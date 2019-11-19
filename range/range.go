package main

import "fmt"

func main(){
	// Array Range: Index, Value
	arr1 := []int{1,2,3,4,5}
	sum := 0
	for i, v := range arr1 {
		fmt.Println("Index:", i)
		fmt.Println("Value:", v)
		sum += v
	}
	fmt.Println("SUM:", sum)

	// Slice Range: Index, Value
	slice1 := make([]string, 4)
	for i := range slice1 {
		slice1[i] = "VALUE"
	}
	for i, v := range slice1 {
		fmt.Println("Index:", i)
		fmt.Println("Value:", v)
	}

	// Map Range: Key, Value
	map1 := map[string]int{"VAL1":10, "VAL2":20, "VAL3":30}
	for k, v := range map1 {
		fmt.Printf("KEY: %s\nVALUE: %d\n", k, v)
	}
	for k1 := range map1 {
		fmt.Println("KEY", k1)
	}
	for _, v := range map1 {
		fmt.Println("VALUE", v)
	}
}