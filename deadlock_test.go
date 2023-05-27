package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Bambang",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Yanto",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 20000)

	time.Sleep(3 * time.Second)

	fmt.Println("name ", user1.Name, "balance ", user1.Balance)
	fmt.Println("name ", user2.Name, "balance ", user2.Balance)

}
