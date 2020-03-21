package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks timer1 channel C unitl it sends value saying timer fired
	<-timer1.C
	fmt.Println("Timer 1 Fired")

	// Benefits of timer vs timeout is simply that you can cancel timers
	timer2 := time.NewTimer(time.Second)
	go func() {
		// Blocks timer2 channel C until timer2 is finished
		<-timer2.C
		fmt.Println("Timer 2 Fired")
	}()
	// Stop timer2
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 Stopped")
	}

	time.Sleep(2 * time.Second)
}