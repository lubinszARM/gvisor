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

package wasm

import (
	//	"fmt"
	//	"os"
	//	"runtime"

	//	"golang.org/x/sys/unix"
	//	"gvisor.dev/gvisor/pkg/log"
	//	"gvisor.dev/gvisor/pkg/procid"
	//"gvisor.dev/gvisor/pkg/sentry/arch"
	"gvisor.dev/gvisor/pkg/sentry/memmap"
	"gvisor.dev/gvisor/pkg/sentry/platform"
	"gvisor.dev/gvisor/pkg/sync"
	"gvisor.dev/gvisor/pkg/usermem"
)

// subprocess is a collection of threads being traced.
type subprocess struct {
	platform.NoAddressSpaceIO

	// requests is used to signal creation of new threads.
	//	requests chan chan *thread

	// mu protects the following fields.
	mu sync.Mutex

	// contexts is the set of contexts for which it's possible that
	// context.lastFaultSP == this subprocess.
	contexts map[*context]struct{}
}

// Release kills the subprocess.
//
// Just kidding! We can't safely co-ordinate the detaching of all the
// tracees (since the tracers are random runtime threads, and the process
// won't exit until tracers have been notifier).
//
// Therefore we simply unmap everything in the subprocess and return it to the
// globalPool. This has the added benefit of reducing creation time for new
// subprocesses.
func (s *subprocess) Release() {
	/*	go func() { // S/R-SAFE: Platform.
			s.unmap()
			globalPool.mu.Lock()
			globalPool.available = append(globalPool.available, s)
			globalPool.mu.Unlock()
		}()
	*/
}

/*
// NotifyInterrupt implements interrupt.Receiver.NotifyInterrupt.
func (t *thread) NotifyInterrupt() {
	//	unix.Tgkill(int(t.tgid), int(t.tid), unix.Signal(platform.SignalInterrupt))
}
*/
// MapFile implements platform.AddressSpace.MapFile.
func (s *subprocess) MapFile(addr usermem.Addr, f memmap.File, fr memmap.FileRange, at usermem.AccessType, precommit bool) error {
	return nil
}

// Unmap implements platform.AddressSpace.Unmap.
func (s *subprocess) Unmap(addr usermem.Addr, length uint64) {
}

// PreFork implements platform.AddressSpace.PreFork.
func (s *subprocess) PreFork() {}

// PostFork implements platform.AddressSpace.PostFork.
func (s *subprocess) PostFork() {}
