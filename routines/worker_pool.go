package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Println("WORKER", id, "STARTED JOB", i)
		time.Sleep(time.Second)
		fmt.Println("WORKER", id, "DONE JOB", i)

		results <- i*2
	}
}

func main() {
	const numOfJobs = 10
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < numOfJobs; j++ {
		jobs <- j
	}

	for r := 0; r < numOfJobs; r++ {
		<-results
	}
}