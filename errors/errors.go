package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg int
	problem string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s", e.problem)
}

func positiveNum(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("Not a positive number")
	}

	return arg, nil
}

func positiveNum2(arg int) (int, error) {
	if arg < 0 {
		return -1, &argError{arg, "Not a positive number"}
	}

	return arg, nil
}

func evenNumber(arg int) (int, error) {
	if arg%2 != 0 {
		return -1, &argError{arg, "Not a even number"}
	} else {
		return arg, nil
	}
}



func main() {
	for _, i := range []int{-1, 1} {
		if n, e := positiveNum(i); e != nil {
			fmt.Println("ERROR", e)
		} else {
			fmt.Println("POSITIVE:", n)
		}
	}

	for _, i := range []int{-100, 23} {
		if n, e := positiveNum2(i); e != nil {
			fmt.Println("ERROR", e)
		} else {
			fmt.Println("POSITIVE:", n)
		}
	}

	for _, i := range []int{23, 25, 26} {
		if n, e := evenNumber(i); e != nil {
			fmt.Println("ERROR", e)
		} else {
			fmt.Println("EVEN NUMBER:", n)
		}
	}
}
