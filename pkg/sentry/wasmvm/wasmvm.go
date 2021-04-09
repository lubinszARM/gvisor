package wasmvm

import (
	"encoding/binary"
	//	"reflect"
	"fmt"

	"gvisor.dev/gvisor/pkg/sentry/syscalls/wasm"
	//"gvisor.dev/gvisor/pkg/log"
	//	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
	//"gvisor.dev/gvisor/pkg/syserror"
	"gvisor.dev/gvisor/pkg/usermem"
)

type WasmVM interface {
	LinkHost(impl WasmVM, thread *WasmThread, t *kernel.Task) error
	LoadWasm(wasmData []byte) error
	RunFunction(functionName string) error
	UnsafeMemory() []byte
}

type WasmVmBase struct {
	impl   WasmVM
	thread *WasmThread
	task   *kernel.Task
}

func (vm *WasmVmBase) LinkHost(impl WasmVM, thread *WasmThread, t *kernel.Task) error {
	vm.impl = impl
	vm.thread = thread
	thread.vm = impl
	vm.task = t

	return nil
}

func (vm *WasmVmBase) WasiFdWrite(t *kernel.Task, fd int32, iovs int32, size int32, written int32) int32 {
	ptr := vm.impl.UnsafeMemory()
	txt := binary.LittleEndian.Uint32(ptr[iovs : iovs+4])
	siz := binary.LittleEndian.Uint32(ptr[iovs+4 : iovs+8])
	fmt.Printf(string(ptr[txt : txt+siz]))
	binary.LittleEndian.PutUint32(ptr[written:written+4], siz)

	ioseq := usermem.BytesIOSequence(ptr[txt : txt+siz])

	/*
		file := t.GetFile(fd)
		if file == nil {
			return 0
		}
		defer file.DecRef(t)

		//	src := usermem.IOSequence{}

		n, err := file.Writev(t, ioseq)
		if err != syserror.ErrWouldBlock || file.Flags().NonBlocking {
			if n > 0 {
				// Queue notification if we wrote anything.
				file.Dirent.InotifyEvent(linux.IN_MODIFY, 0)
			}
			//return n, err
			return 0
		}
	*/
	return wasm.FdWrite(t, fd, ioseq, written)
	return 0
	//return int32(siz)
}
