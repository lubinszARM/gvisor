package wasmvm

import (
	pkgcontext "gvisor.dev/gvisor/pkg/context"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
)

type WasmProc struct {
	WasmThread
}

var GoWasmVM WasmVM

func NewWasmProc(vm WasmVM, ctx pkgcontext.Context) (*WasmProc, error) {
	thread := &WasmProc{}
	if GoWasmVM != nil {
		vm = GoWasmVM
	}

	task := ctx.(*kernel.Task)
	err := thread.InitVM(vm, task)
	if err != nil {
		return nil, err
	}
	return thread, nil
}
