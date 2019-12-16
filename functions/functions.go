package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func add4(a, b, c, d int) int {
	return a + b + c + d
}

func addAndSub(a, b int) (int, int) {
	return a+b, a-b
}

func addInfinite(nums...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func count() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("1+1 =", add(1,2))
	fmt.Println("1+2+3+4 =", add4(1,2,3,4))
	addResult, subResult := addAndSub(1,1)
	fmt.Printf("1+1 = %d\n1-1 = %d\n", addResult, subResult)
	// Veriadic function: unlimited unmber of args
	largeSum := addInfinite(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("1+2+3+...+10 =", largeSum)
	// Pass in slice/array with ...
	largeSum = addInfinite([]int{11,12,13,14,15,16,17,18,19,20}...)
	fmt.Println("11+12+...+20 =", largeSum)
	// Closure
	countNum := count()
	fmt.Println("START COUNT")
	fmt.Println(countNum())
	fmt.Println(countNum())
	fmt.Println(countNum())
	fmt.Println("END COUNT")
}
