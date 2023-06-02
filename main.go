package main

import (
	"context"
	_ "embed"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

//go:embed consumer.wasm
var _WASM_Consumer []byte

//go:embed server.wasm
var _WASM_Server []byte

var _ExecWasmToWazero func()

func init() {
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	host := r.NewHostModuleBuilder("env")
	{
		fn := host.NewFunctionBuilder()
		counter := uint64(0)
		fn.WithGoFunction(api.GoFunc(func(ctx context.Context, stack []uint64) {
			counter++
			stack[0] = counter
		}), nil, []api.ValueType{api.ValueTypeI64})
		fn.Export("somefunction")
	}

	_, err := host.Instantiate(ctx)
	if err != nil {
		panic(err)
	}

	instantiate, err := r.Instantiate(ctx, _WASM_Consumer)
	if err != nil {
		panic(err)
	}

	_ExecWasmToWazero = func() {
		if _, err := instantiate.ExportedFunction("work").Call(ctx); err != nil {
			panic(err)
		}
	}
}

var _ExecWasmToWasm func()

func init() {
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)

	_, err := r.InstantiateWithConfig(ctx, _WASM_Server, wazero.NewModuleConfig().WithName("env"))
	if err != nil {
		panic(err)
	}

	instantiate, err := r.Instantiate(ctx, _WASM_Consumer)
	if err != nil {
		panic(err)
	}

	_ExecWasmToWasm = func() {
		if _, err := instantiate.ExportedFunction("work").Call(ctx); err != nil {
			panic(err)
		}
	}
}

var _ExecWazeroToWasm func()

func init() {
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)

	module, err := r.InstantiateWithConfig(ctx, _WASM_Server, wazero.NewModuleConfig().WithName("env"))
	if err != nil {
		panic(err)
	}

	counter := module.ExportedFunction("somefunction")

	_ExecWazeroToWasm = func() {
		for i := 0; i < 100; i++ {
			_, err = counter.Call(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}
