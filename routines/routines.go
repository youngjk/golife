package main

import (
	"fmt"
	"time"
)

func print(mode string) {
	for i := 0; i < 3; i++ {
		fmt.Println("MODE:", mode, ":", i)
	}
}

func main() {
	print("SYNC")
	go print("ASYNC")

	go fmt.Println("HELLO WORLD")
	time.Sleep(time.Second)
  fmt.Println("done")
}
