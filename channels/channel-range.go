package main

import "fmt"

func main() {
	queue := make(chan string, 3)
	queue <- "ONE"
	queue <- "TWO"
	queue <- "THREE"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
