package wasmvm

import (
	"errors"

	wasmtime "github.com/bytecodealliance/wasmtime-go"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
	//	"gvisor.dev/gvisor/pkg/sentry/syscalls/wasm"
)

type WasmTimeVM struct {
	WasmVmBase
	instance *wasmtime.Instance
	linker   *wasmtime.Linker
	memory   *wasmtime.Memory
	module   *wasmtime.Module
	store    *wasmtime.Store
}

func NewWasmTimeVM() *WasmTimeVM {
	vm := &WasmTimeVM{}
	vm.store = wasmtime.NewStore(wasmtime.NewEngine())
	vm.linker = wasmtime.NewLinker(vm.store)
	return vm
}

func (vm *WasmTimeVM) LinkHost(impl WasmVM, thread *WasmThread, t *kernel.Task) error {
	vm.WasmVmBase.LinkHost(impl, thread, t)

	// go implementation uses this one to write panic message
	err := vm.linker.DefineFunc("wasi_unstable", "fd_write",
		func(fd int32, iovs int32, size int32, written int32) int32 {
			return vm.WasiFdWrite(t, fd, iovs, size, written)
		})
	if err != nil {
		return err
	}

	err = vm.linker.DefineFunc("wasi_snapshot_preview1", "fd_write",
		func(fd int32, iovs int32, size int32, written int32) int32 {
			return vm.WasiFdWrite(t, fd, iovs, size, written)
		})
	if err != nil {
		return err
	}
	return nil
}

func (vm *WasmTimeVM) LoadWasm(wasmData []byte) error {
	var err error
	vm.module, err = wasmtime.NewModule(vm.store.Engine, wasmData)
	if err != nil {
		return err
	}
	vm.instance, err = vm.linker.Instantiate(vm.module)
	if err != nil {
		return err
	}
	memory := vm.instance.GetExport("memory")
	if memory == nil {
		return errors.New("no memory export")
	}
	vm.memory = memory.Memory()
	if vm.memory == nil {
		return errors.New("not a memory type")
	}
	return nil
}

func (vm *WasmTimeVM) RunFunction(functionName string) error {
	export := vm.instance.GetExport(functionName)
	if export == nil {
		return errors.New("unknown export function: '" + functionName + "'")
	}
	_, err := export.Func().Call()
	return err
}

func (vm *WasmTimeVM) UnsafeMemory() []byte {
	return vm.memory.UnsafeData()
}
