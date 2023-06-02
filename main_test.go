package main

import (
	"testing"
)

func BenchmarkNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ExecWasmToWazero()
	}
}

func BenchmarkWasm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ExecWasmToWasm()
	}
}

func BenchmarkWazero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ExecWazeroToWasm()
	}
}
