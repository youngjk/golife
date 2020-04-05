package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// unsigned integer to use as our counter (Always positive)
	var ops uint64
	// WaitGroup will help wait for all goroutines to finish
	var wg sync.WaitGroup

	// We'll start 50 goroutines which increments counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Wait until goroutines are done
	wg.Wait()
	// Print counter at termination 50 * 1000 = 50,000
	fmt.Println("OPS:", ops)
}