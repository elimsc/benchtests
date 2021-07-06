package benchtests_test

import (
	"math"
	"testing"
)

func BenchmarkWay1(b *testing.B) {
	x, y := math.MaxInt32, math.MaxInt32
	for i := 0; i < b.N; i++ {
		_ = x + (y-x)/2
	}
}

func BenchmarkWay2(b *testing.B) {
	x, y := math.MaxInt32, math.MaxInt32
	for i := 0; i < b.N; i++ {
		_ = int(uint(x+y) / 2)
	}
}

// 结论: 第二个快一点点(真的只有一点点)
