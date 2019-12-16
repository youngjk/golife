package main

import "fmt"

func addPowerThree(num int) int {
	fmt.Println(num)
	if num <= 0 {
		return 0
	} else if num%3 != 0 {
		return num + addPowerThree(num - num%3)
	}
	return num + addPowerThree(num-3)
}

func main() {
	fmt.Println("Sum of mulitple of 3:", addPowerThree(100))
}
