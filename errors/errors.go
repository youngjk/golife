package main

import (
	"errors"
	"fmt"
)

func positiveNum(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("Not a positive number")
	}

	return arg, nil
}

func main() {
	for _, i := range []int{-1, 1} {
		if n, e := positiveNum(i); e != nil {
			fmt.Println("ERROR", e)
		} else {
			fmt.Println("POSITIVE:", n)
		}
	}
}
