package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// On return notify waiting group we are done
	defer wg.Done()

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// Waitgroup waits for all go routines launched here to finish
	var wg sync.WaitGroup

	// Launch several workers and increment waitgroup counter for each worker
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until WaitGroup counter goes back to 0
	wg.Wait()
}