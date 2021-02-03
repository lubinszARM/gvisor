// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stack

import (
	"fmt"
	"time"

	"gvisor.dev/gvisor/pkg/sync"
	"gvisor.dev/gvisor/pkg/tcpip"
)

const linkAddrCacheSize = 512 // max cache entries

// linkAddrCache is a fixed-sized cache mapping IP addresses to link addresses.
//
// The entries are stored in a ring buffer, oldest entry replaced first.
//
// This struct is safe for concurrent use.
type linkAddrCache struct {
	nic *NIC

	linkRes LinkAddressResolver

	// ageLimit is how long a cache entry is valid for.
	ageLimit time.Duration

	// resolutionTimeout is the amount of time to wait for a link request to
	// resolve an address.
	resolutionTimeout time.Duration

	// resolutionAttempts is the number of times an address is attempted to be
	// resolved before failing.
	resolutionAttempts int

	mu struct {
		sync.Mutex
		table map[tcpip.Address]*linkAddrEntry
		lru   linkAddrEntryList
	}
}

// entryState controls the state of a single entry in the cache.
type entryState int

const (
	// incomplete means that there is an outstanding request to resolve the
	// address. This is the initial state.
	incomplete entryState = iota
	// ready means that the address has been resolved and can be used.
	ready
)

// String implements Stringer.
func (s entryState) String() string {
	switch s {
	case incomplete:
		return "incomplete"
	case ready:
		return "ready"
	default:
		return fmt.Sprintf("unknown(%d)", s)
	}
}

// A linkAddrEntry is an entry in the linkAddrCache.
// This struct is thread-compatible.
type linkAddrEntry struct {
	// linkAddrEntryEntry access is synchronized by the linkAddrCache lock.
	linkAddrEntryEntry

	cache *linkAddrCache

	mu struct {
		sync.RWMutex

		addr       tcpip.Address
		linkAddr   tcpip.LinkAddress
		expiration time.Time
		s          entryState

		// done is closed when address resolution is complete. It is nil iff s is
		// incomplete and resolution is not yet in progress.
		done chan struct{}

		// onResolve is called with the result of address resolution.
		onResolve []func(LinkResolutionResult)
	}
}

func (e *linkAddrEntry) notifyCompletionLocked(linkAddr tcpip.LinkAddress) {
	res := LinkResolutionResult{LinkAddress: linkAddr, Success: len(linkAddr) != 0}
	for _, callback := range e.mu.onResolve {
		callback(res)
	}
	e.mu.onResolve = nil
	if ch := e.mu.done; ch != nil {
		close(ch)
		e.mu.done = nil
		// Dequeue the pending packets in a new goroutine to not hold up the current
		// goroutine as writing packets may be a costly operation.
		//
		// At the time of writing, when writing packets, a neighbor's link address
		// is resolved (which ends up obtaining the entry's lock) while holding the
		// link resolution queue's lock. Dequeuing packets in a new goroutine avoids
		// a lock ordering violation.
		go e.cache.nic.linkResQueue.dequeue(ch, linkAddr, len(linkAddr) != 0)
	}
}

// changeStateLocked sets the entry's state to ns.
//
// The entry's expiration is bumped up to the greater of itself and the passed
// expiration; the zero value indicates immediate expiration, and is set
// unconditionally - this is an implementation detail that allows for entries
// to be reused.
//
// Precondition: e.mu must be locked
func (e *linkAddrEntry) changeStateLocked(ns entryState, expiration time.Time) {
	if e.mu.s == incomplete && ns == ready {
		e.notifyCompletionLocked(e.mu.linkAddr)
	}

	if expiration.IsZero() || expiration.After(e.mu.expiration) {
		e.mu.expiration = expiration
	}
	e.mu.s = ns
}

// add adds a k -> v mapping to the cache.
func (c *linkAddrCache) add(k tcpip.Address, v tcpip.LinkAddress) {
	// Calculate expiration time before acquiring the lock, since expiration is
	// relative to the time when information was learned, rather than when it
	// happened to be inserted into the cache.
	expiration := time.Now().Add(c.ageLimit)

	c.mu.Lock()
	entry := c.getOrCreateEntryLocked(k)
	entry.mu.Lock()
	defer entry.mu.Unlock()
	c.mu.Unlock()

	entry.mu.linkAddr = v
	entry.changeStateLocked(ready, expiration)
}

// getOrCreateEntryLocked retrieves a cache entry associated with k. The
// returned entry is always refreshed in the cache (it is reachable via the
// map, and its place is bumped in LRU).
//
// If a matching entry exists in the cache, it is returned. If no matching
// entry exists and the cache is full, an existing entry is evicted via LRU,
// reset to state incomplete, and returned. If no matching entry exists and the
// cache is not full, a new entry with state incomplete is allocated and
// returned.
func (c *linkAddrCache) getOrCreateEntryLocked(k tcpip.Address) *linkAddrEntry {
	if entry, ok := c.mu.table[k]; ok {
		c.mu.lru.Remove(entry)
		c.mu.lru.PushFront(entry)
		return entry
	}
	var entry *linkAddrEntry
	if len(c.mu.table) == linkAddrCacheSize {
		entry = c.mu.lru.Back()
		entry.mu.Lock()

		delete(c.mu.table, entry.mu.addr)
		c.mu.lru.Remove(entry)

		// Wake waiters and mark the soon-to-be-reused entry as expired.
		entry.notifyCompletionLocked("" /* linkAddr */)
		entry.mu.Unlock()
	} else {
		entry = new(linkAddrEntry)
	}

	*entry = linkAddrEntry{
		cache: c,
	}
	entry.mu.Lock()
	entry.mu.addr = k
	entry.mu.s = incomplete
	entry.mu.Unlock()
	c.mu.table[k] = entry
	c.mu.lru.PushFront(entry)
	return entry
}

// get reports any known link address for addr.
func (c *linkAddrCache) get(addr, localAddr tcpip.Address, onResolve func(LinkResolutionResult)) (tcpip.LinkAddress, <-chan struct{}, tcpip.Error) {
	c.mu.Lock()
	entry := c.getOrCreateEntryLocked(addr)
	entry.mu.Lock()
	defer entry.mu.Unlock()
	c.mu.Unlock()

	switch s := entry.mu.s; s {
	case ready:
		if !time.Now().After(entry.mu.expiration) {
			// Not expired.
			if onResolve != nil {
				onResolve(LinkResolutionResult{LinkAddress: entry.mu.linkAddr, Success: true})
			}
			return entry.mu.linkAddr, nil, nil
		}

		entry.changeStateLocked(incomplete, time.Time{})
		fallthrough
	case incomplete:
		if onResolve != nil {
			entry.mu.onResolve = append(entry.mu.onResolve, onResolve)
		}
		if entry.mu.done == nil {
			entry.mu.done = make(chan struct{})
			go c.startAddressResolution(addr, localAddr, entry.mu.done) // S/R-SAFE: link non-savable; wakers dropped synchronously.
		}
		return entry.mu.linkAddr, entry.mu.done, &tcpip.ErrWouldBlock{}
	default:
		panic(fmt.Sprintf("invalid cache entry state: %s", s))
	}
}

func (c *linkAddrCache) startAddressResolution(k tcpip.Address, localAddr tcpip.Address, done <-chan struct{}) {
	for i := 0; ; i++ {
		// Send link request, then wait for the timeout limit and check
		// whether the request succeeded.
		c.linkRes.LinkAddressRequest(k, localAddr, "" /* linkAddr */)

		select {
		case now := <-time.After(c.resolutionTimeout):
			if stop := c.checkLinkRequest(now, k, i); stop {
				return
			}
		case <-done:
			return
		}
	}
}

// checkLinkRequest checks whether previous attempt to resolve address has
// succeeded and mark the entry accordingly. Returns true if request can stop,
// false if another request should be sent.
func (c *linkAddrCache) checkLinkRequest(now time.Time, k tcpip.Address, attempt int) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.mu.table[k]
	if !ok {
		// Entry was evicted from the cache.
		return true
	}
	entry.mu.Lock()
	defer entry.mu.Unlock()

	switch s := entry.mu.s; s {
	case ready:
		// Entry was made ready by resolver.
	case incomplete:
		if attempt+1 < c.resolutionAttempts {
			// No response yet, need to send another ARP request.
			return false
		}
		// Max number of retries reached, delete entry.
		entry.notifyCompletionLocked("" /* linkAddr */)
		delete(c.mu.table, k)
	default:
		panic(fmt.Sprintf("invalid cache entry state: %s", s))
	}
	return true
}

func (c *linkAddrCache) init(nic *NIC, ageLimit, resolutionTimeout time.Duration, resolutionAttempts int, linkRes LinkAddressResolver) {
	*c = linkAddrCache{
		nic:                nic,
		linkRes:            linkRes,
		ageLimit:           ageLimit,
		resolutionTimeout:  resolutionTimeout,
		resolutionAttempts: resolutionAttempts,
	}

	c.mu.Lock()
	c.mu.table = make(map[tcpip.Address]*linkAddrEntry, linkAddrCacheSize)
	c.mu.Unlock()
}

var _ neighborTable = (*linkAddrCache)(nil)

func (*linkAddrCache) neighbors() ([]NeighborEntry, tcpip.Error) {
	return nil, &tcpip.ErrNotSupported{}
}

func (c *linkAddrCache) addStaticEntry(addr tcpip.Address, linkAddr tcpip.LinkAddress) {
	c.add(addr, linkAddr)
}

func (*linkAddrCache) remove(addr tcpip.Address) tcpip.Error {
	return &tcpip.ErrNotSupported{}
}

func (*linkAddrCache) removeAll() tcpip.Error {
	return &tcpip.ErrNotSupported{}
}

func (c *linkAddrCache) handleProbe(addr tcpip.Address, linkAddr tcpip.LinkAddress) {
	if len(linkAddr) != 0 {
		// NUD allows probes without a link address but linkAddrCache
		// is a simple neighbor table which does not implement NUD.
		//
		// As per RFC 4861 section 4.3,
		//
		//   Source link-layer address
		//               The link-layer address for the sender. MUST NOT be
		//               included when the source IP address is the
		//               unspecified address. Otherwise, on link layers
		//               that have addresses this option MUST be included in
		//               multicast solicitations and SHOULD be included in
		//               unicast solicitations.
		c.add(addr, linkAddr)
	}
}

func (c *linkAddrCache) handleConfirmation(addr tcpip.Address, linkAddr tcpip.LinkAddress, flags ReachabilityConfirmationFlags) {
	if len(linkAddr) != 0 {
		// NUD allows confirmations without a link address but linkAddrCache
		// is a simple neighbor table which does not implement NUD.
		//
		// As per RFC 4861 section 4.4,
		//
		//   Target link-layer address
		//               The link-layer address for the target, i.e., the
		//               sender of the advertisement. This option MUST be
		//               included on link layers that have addresses when
		//               responding to multicast solicitations. When
		//               responding to a unicast Neighbor Solicitation this
		//               option SHOULD be included.
		c.add(addr, linkAddr)
	}
}

func (c *linkAddrCache) handleUpperLevelConfirmation(tcpip.Address) {}

func (*linkAddrCache) nudConfig() (NUDConfigurations, tcpip.Error) {
	return NUDConfigurations{}, &tcpip.ErrNotSupported{}
}

func (*linkAddrCache) setNUDConfig(NUDConfigurations) tcpip.Error {
	return &tcpip.ErrNotSupported{}
}
