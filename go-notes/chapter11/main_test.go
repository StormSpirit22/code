package main

import (
	"testing"
)
func heap() []byte {
	return make([]byte, 1024*10)
}

func BenchmarkHeap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = heap()
	}
}