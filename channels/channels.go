package main

import (
	"fmt"
	"strconv"
	"time"
)

func task(done chan bool) {
	fmt.Println("Starting Task...")
	time.Sleep(time.Second)
	fmt.Println("Done Task")

	done <- true
}

// Channel sender/receiver behaviour specification
func senderTunnel(pings chan<- string, pingString string) {
	pings <- pingString
}

func receiveTunnel(pongs <-chan string) string {
	return <-pongs
}

func main() {
	pingChannel := make(chan string)
	pongChannel := make(chan string)

	go func() {
		pingChannel <- "PING"
	}()

	go func() {
		pongChannel <- "PONG"
	}()

	pingResult := <-pingChannel
	pongResult := <-pongChannel
	fmt.Println(pingResult)
	fmt.Println(pongResult)

	// Channels only receive limited number of values without corresponding receiver
	const CHANNEL_COUNT = 5

	fiveCapChannel := make(chan string, CHANNEL_COUNT)
	for i := 1; i <= CHANNEL_COUNT; i++ {
		fiveCapChannel <- strconv.Itoa(i)
	}

	for i := 0; i < CHANNEL_COUNT; i++ {
		fmt.Println(<-fiveCapChannel)
	}

	// Channel Synchronization use
	doneChannel := make(chan bool, 1)
	go task(doneChannel)

	if <-doneChannel {
		fmt.Println("Channel Signal Received")
	}

	// Channel restricted params
	ch1 := make(chan string, 1)

	go senderTunnel(ch1, "CHANNEL1 INPUT")
	time.Sleep(time.Second)
	output := receiveTunnel(ch1)
	fmt.Println("RECEIVED:", output)

	// Channel Select
	chSelect1 := make(chan string)
	chSelect2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chSelect1 <- "SELECT1"
	}()
	go func() {
		time.Sleep(4 * time.Second)
		chSelect2 <- "SELECT2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chSelect1:
			fmt.Println("RECEIVED:", msg1)
		case msg2 := <-chSelect2:
			fmt.Println("RECEIVED:", msg2)
		}
	}

	// Channel Timeouts
	chTimeout1 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chTimeout1 <- "5SECONDS"
	}()

	select {
	case chTimeoutMsg1 := <-chTimeout1:
		fmt.Println("RECEIVED:", chTimeoutMsg1)
	case <-time.After(1 * time.Second):
		fmt.Println("1 SECOND TIMEOUT")
	}

	chTimeout2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chTimeout2 <- "2SECONDS"
	}()

	select {
	case chTimeoutMsg2 := <-chTimeout2:
		fmt.Println("RECEIVED:", chTimeoutMsg2)
	case <-time.After(5 * time.Second):
		fmt.Println("5 SECOND TIMEOUT")
	}
}
