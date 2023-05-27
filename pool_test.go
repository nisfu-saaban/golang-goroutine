package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool sync.Pool

	pool.Put("Andi")
	pool.Put("Kick")
	pool.Put("Ass")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("ambil data ", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("selesai")
}
