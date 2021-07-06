package benchtests_test

import (
	"fmt"
	"testing"
)

func BenchmarkStringAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "hello " + "world"
	}
}

func BenchmarkStringFprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello %s", "world")
	}
}

// 差距很大: 0.2779 ns/op vs 79.04 ns/op
