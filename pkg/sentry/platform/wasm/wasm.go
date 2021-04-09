// Copyright 2021 The gVisor Authors.
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
	"os"

	pkgcontext "gvisor.dev/gvisor/pkg/context"
	"gvisor.dev/gvisor/pkg/log"
	"gvisor.dev/gvisor/pkg/sentry/arch"
	"gvisor.dev/gvisor/pkg/sentry/loader"
	"gvisor.dev/gvisor/pkg/sentry/platform"
	"gvisor.dev/gvisor/pkg/sentry/platform/interrupt"
	"gvisor.dev/gvisor/pkg/sentry/wasmvm"
	"gvisor.dev/gvisor/pkg/sync"
	"gvisor.dev/gvisor/pkg/usermem"
)

type context struct {
	// interrupt is the interrupt context.
	interrupt interrupt.Forwarder

	// mu protects the following fields.
	mu sync.Mutex
}

func setExit(regs *arch.Registers) {
	// exit_group = 231
	regs.Orig_rax = 231
	regs.Rax = 231
}

// Switch runs the provided context in the given address space.
func (c *context) Switch(ctx pkgcontext.Context, mm platform.MemoryManager, ac arch.Context, cpu int32) (*arch.SignalInfo, usermem.AccessType, error) {
	errChan := make(chan error)
	go func() {
		glbWASI, err := wasmvm.NewWasmProc(wasmvm.NewWasmTimeVM(), ctx)
		if err != nil {
			log.Infof("WASM NewWasmProcessor fail:%v.\n", err)
			//return loadedELF{}, nil, nil, nil, syserror.ENOEXEC
		}

		glbWASI.Init()
		err = glbWASI.LoadWasm(loader.GBuf)
		if err != nil {
			log.Infof("WASM loadwasm fail:%v.\n", err)
		}

		glbWASI.RunFunction("_start")

		errChan <- nil
	}()

	// Wait until error or readiness.
	if err := <-errChan; err != nil {
		log.Infof("WASM failed:%v.\n", err)
	}

	setExit(&ac.StateData().Regs)

	return nil, usermem.NoAccess, nil
}

// Interrupt interrupts the running guest application associated with this context.
func (c *context) Interrupt() {
	c.interrupt.NotifyInterrupt()
}

// Release implements platform.Context.Release().
func (c *context) Release() {}

// FullStateChanged implements platform.Context.FullStateChanged.
func (c *context) FullStateChanged() {}

// PullFullState implements platform.Context.PullFullState.
func (c *context) PullFullState(as platform.AddressSpace, ac arch.Context) {}

// WASM represents a collection of wasm subprocesses.
type WASM struct {
	platform.MMapMinAddr
	platform.NoCPUPreemptionDetection
	platform.UseHostGlobalMemoryBarrier
}

// New returns a new wasm-based implementation of the platform interface.
func New() (*WASM, error) {
	return &WASM{}, nil
}

// SupportsAddressSpaceIO implements platform.Platform.SupportsAddressSpaceIO.
func (*WASM) SupportsAddressSpaceIO() bool {
	return false
}

// CooperativelySchedulesAddressSpace implements platform.Platform.CooperativelySchedulesAddressSpace.
func (*WASM) CooperativelySchedulesAddressSpace() bool {
	return false
}

// MapUnit implements platform.Platform.MapUnit.
func (*WASM) MapUnit() uint64 {
	// The host kernel manages page tables and arbitrary-sized mappings
	// have effectively the same cost.
	return 0
}

// MaxUserAddress returns the first address that may not be used by user
// applications.
func (*WASM) MaxUserAddress() usermem.Addr {
	//return usermem.Addr(stubStart)
	return 0xffff000000000000
}

// NewAddressSpace returns a new subprocess.
func (p *WASM) NewAddressSpace(_ interface{}) (platform.AddressSpace, <-chan struct{}, error) {
	//	as, err := newSubprocess(globalPool.master.createStub)
	// Return the new address space.
	return &subprocess{}, nil, nil
}

// NewContext returns an interruptible context.
func (*WASM) NewContext() platform.Context {
	return &context{}
}

type constructor struct{}

func (*constructor) New(*os.File) (platform.Platform, error) {
	return New()
}

func (*constructor) OpenDevice() (*os.File, error) {
	return nil, nil
}

// Flags implements platform.Constructor.Flags().
func (*constructor) Requirements() platform.Requirements {
	return platform.Requirements{}
}

func init() {
	platform.Register("wasm", &constructor{})
}
