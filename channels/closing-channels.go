package main

import (
	"fmt"
)

func main() {
	threads := make(chan int, 5)
	done :=  make(chan bool)

	// Infinite loop until all threads are received
	go func() {
		for {
			t, more := <- threads
			if more {
				fmt.Println("Received Thread", t)
			} else {
				fmt.Println("Received all Threads")
				done <- true
				return
			}
		}
	}()

	for t := 1; t <= 3; t++ {
		threads <- t
		fmt.Println("SENT THREADS", t)
	}
	close(threads)
	fmt.Println("SENT ALL THREADS")

	<-done
}