package golanggoroutine

import (
	"fmt"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	channel := make(chan any, 3)
	defer close(channel)

	go func() {
		channel <- "Suharto"
		channel <- "Raya"
		channel <- "Hambu"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		// fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")

}

func TestRangeChannel(t *testing.T) {
	channel := make(chan any)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Meterima data ke ", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan any)
	channel2 := make(chan any)

	defer close(channel1)
	defer close(channel2)

	count := 0

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	for {
		select {
		case data := <-channel1:
			fmt.Println("channel 1 ", data)
			count++
		case data := <-channel2:
			fmt.Println("channel 2 ", data)
			count++
		}

		if count == 2 {
			break
		}
	}
}
