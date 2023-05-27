package golanggoroutine

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int, tastQueue chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tastQueue {
		fmt.Printf("worker %d started task %d\n", id, task)
		time.Sleep(1 * time.Second)
		fmt.Printf("worker %d complete task %d\n", id, task)
	}
}
