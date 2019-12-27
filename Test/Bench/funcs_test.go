package randomrune

import (
	"testing"
)

// BenchmarkRandomIndexSlice
func BenchmarkRandomIndexSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomIndexSlice(10000)
	}
}

// BenchmarkRandomASCII
func BenchmarkRandomASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomASCII(10000)
	}
}
