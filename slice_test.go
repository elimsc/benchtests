package benchtests_test

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	// 0.2750 ns/op
	for i := 0; i < b.N; i++ {
		a := []int{1, 2, 3, 0, 0, 0}
		b := []int{4, 5, 6}
		_ = append(a[:3], b...)
	}
}

func BenchmarkSliceCopy(b *testing.B) {
	// 0.2750 ns/op
	for i := 0; i < b.N; i++ {
		a := []int{1, 2, 3, 0, 0, 0}
		b := []int{4, 5, 6}
		copy(a[3:], b)
	}
}
