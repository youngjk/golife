package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync/atomic"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	// Counts of read and write operations
	var readOps uint64
	var writeOps uint64

	// Read and Write channels will be used by goroutines to issue read and write requests
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Goroutine owning state: Map private to stateful goroutine
	// Repeatedly selects read and rights channels to respond to requests
	// Response is executed by performing a request and sending result to response channel
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// On this go routine we generate 100 read requests, we send to channel and wait for response right away
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// On this go routine we generate 10 write requests, we send to channel and wait for response right away
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Execute for 1 second
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("READOPS COUNT:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("WRITEOPS COUNT:", writeOpsFinal)
}