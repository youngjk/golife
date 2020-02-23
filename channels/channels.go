package main

import (
	"fmt"
	"strconv"
)

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
	const CHANNEL_COUNT := 5

	fiveCapChannel := make(chan string, CHANNEL_COUNT)
	for i := 1; i <= CHANNEL_COUNT; i++ {
		fiveCapChannel <- strconv.Itoa(i)
	}

	for i := 0; i < CHANNEL_COUNT; i++ {
		fmt.Println(<-fiveCapChannel)
	}
}