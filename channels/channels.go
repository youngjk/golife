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
}