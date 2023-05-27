package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	numWorkers := 5
	taskQueue := make(chan int, 10)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(i, taskQueue, &wg)
	}

	for j := 0; j < 10; j++ {
		taskQueue <- j
	}

	close(taskQueue)
	wg.Wait()
	fmt.Println("All Task Complete")
}
