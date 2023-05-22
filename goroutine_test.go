package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestRunHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Setelah")

	time.Sleep(1 * time.Second)
}

func TestChannel(t *testing.T) {
	channel := make(chan any)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Edi Rahmayadi"
		fmt.Println("selesai mengirim data")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

func GiveMeResponse(channel chan any) {
	time.Sleep(2 * time.Second)
	channel <- "Abdurrahman Wahab"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan any)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- any) {
	time.Sleep(2 * time.Second)
	channel <- "Abdurrahman Wahab"
}

func OnlyOut(channel <-chan any) {
	data := <-channel
	fmt.Println(data)
}

func TestInAndOutChannel(t *testing.T) {

	channel := make(chan any)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}
