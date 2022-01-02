package benchtests_test

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	// 34.58 ns/op
	for i := 0; i < b.N; i++ {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		_ = append(a, b...)
	}
}

func BenchmarkSliceCopy(b *testing.B) {
	// 1.381 ns/op
	for i := 0; i < b.N; i++ {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		copy(a[len(a):], b)
	}
}
