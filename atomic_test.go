package golanggoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			for i := 0; i < 10; i++ {
				atomic.AddInt64(&x, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count ", x)
}
