package main

import (
	"fmt"
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
}