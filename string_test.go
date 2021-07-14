package benchtests_test

import (
	"fmt"
	"testing"
	"unsafe"
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

func BenchmarkSliceStringConvert1(b *testing.B) {
	s := []byte("hello")
	for i := 0; i < b.N; i++ {
		_ = string(s)
	}
}

func BenchmarkSliceStringConvert2(b *testing.B) {
	s := []byte("hello")
	for i := 0; i < b.N; i++ {
		_ = *(*string)(unsafe.Pointer(&s))
	}
}

// 4.245 ns/op vs 0.2740 ns/op

func BenchmarkStringSliceConvert1(b *testing.B) {
	s := "hello"
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkStringSliceConvert2(b *testing.B) {
	s := "hello"
	for i := 0; i < b.N; i++ {
		_ = *(*([]byte))(unsafe.Pointer(&s))
	}
}

// 5.710 ns/op vs 0.2732 ns/op
