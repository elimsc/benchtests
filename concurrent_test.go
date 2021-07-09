package benchtests_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkConcurrentMutexAdd(b *testing.B) {
	var a int32
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			a += 2
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkConcurrentAtomicAdd(b *testing.B) {
	var a int32
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&a, 2)
			wg.Done()
		}()
	}
	wg.Wait()
}

// 差距很小
