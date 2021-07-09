package benchtests_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

// https://medium.com/@genchilu/whats-false-sharing-and-how-to-solve-it-using-golang-as-example-ef978a305e10

type MyAtomic interface {
	IncreaseAllEles()
	// IncreaseA()
	// IncreaseB()
}

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (myatomic *NoPad) IncreaseAllEles() {
	atomic.AddUint64(&myatomic.a, 1)
	atomic.AddUint64(&myatomic.b, 1)
	atomic.AddUint64(&myatomic.c, 1)
}

type Pad struct {
	a   uint64
	_p1 [8]uint64
	b   uint64
	_p2 [8]uint64
	c   uint64
	_p3 [8]uint64
}

func (myatomic *Pad) IncreaseAllEles() {
	atomic.AddUint64(&myatomic.a, 1)
	atomic.AddUint64(&myatomic.b, 1)
	atomic.AddUint64(&myatomic.c, 1)
}

func testAtomicIncrease(myatomic MyAtomic) {
	paraNum := 1000
	addTimes := 1000
	var wg sync.WaitGroup
	wg.Add(paraNum)
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				myatomic.IncreaseAllEles()
			}
			wg.Done()
		}()
	}
	wg.Wait()

}

func BenchmarkFalseSharingNoPad(b *testing.B) {
	myatomic := &NoPad{}
	b.ResetTimer()
	testAtomicIncrease(myatomic)
}

func BenchmarkFalseSharingPad(b *testing.B) {
	myatomic := &Pad{}
	b.ResetTimer()
	testAtomicIncrease(myatomic)
}

// 0.42 vs 0.28, 按百分比来说，提升还是挺明显的
