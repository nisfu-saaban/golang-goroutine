package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var wg = &sync.WaitGroup{}

func WaitCondition(value int) {
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done ", value)
	cond.L.Unlock()
	wg.Done()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	wg.Wait()
}
