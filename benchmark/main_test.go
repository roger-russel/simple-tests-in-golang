package main

import "testing"

func Benchmark_fibonacci(b *testing.B) {
	f := fibonacci()
	for n := 0; n < b.N; n++ {
		f()
	}
}
