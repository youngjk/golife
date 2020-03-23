package main

import (
	"fmt"
	"time"
)

func main() {
	// Tickers are for regular interval countings
	ticker := time.NewTicker(500 * time.Millisecond)
	fmt.Println("TICKER STARTING")
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("TICKED:", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("TICKER DONE")
}