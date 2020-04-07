package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// State
	var state = make(map[int]int)
	// Mutext synchronize access to state
	var mutex = &sync.Mutex{}
	// Keep track of read/writes we do
	var readOps uint64
	var writeOps uint64

	// Start 100 goroutines to execute repeated reads against state, once per millisecond in each goroutine
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// On read make sure to lock for unique permission to state
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				// Increment readOps Count
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 goroutines to simulate writes, using the same patter as reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// On write make sure to lock
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				// Increase writeOps Count
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("READ OPS COUNT:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("WRITE OPS COUNT:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("STATE:", state)
	mutex.Unlock()
}