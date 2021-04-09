package wasmvm

import (
	"gvisor.dev/gvisor/pkg/sentry/kernel"
)

type WasmThread struct {
	vm WasmVM
}

func (thread *WasmThread) InitVM(vm WasmVM, t *kernel.Task) error {
	return vm.LinkHost(vm, thread, t)
}

//func (host *WasmHost) Init(null HostObject, root HostObject) {
func (thread *WasmThread) Init() {
}

func (thread *WasmThread) LoadWasm(wasmData []byte) error {
	err := thread.vm.LoadWasm(wasmData)
	if err != nil {
		return err
	}
	return nil
}

func (thread *WasmThread) RunFunction(functionName string) (err error) {
	thread.vm.RunFunction(functionName)
	return nil
}
