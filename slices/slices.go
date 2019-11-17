package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Slice: Dynamic Arrays
	slice := make([]string, 3)
	// [  ] NOTE: one less index when empty
	fmt.Println("Empty Slices:", slice, "Size:", len(slice))

	for i := 0; i < len(slice); i++ {
		// GoLang is very type sensitive
		slice[i] = strconv.Itoa(i+1)
	}
	// [1, 2, 3]
	fmt.Println("Updated Slices:", slice, "Size:", len(slice))

	// Append: Adds new value to slice
	for i := len(slice); i < 6; i++ {
		slice = append(slice, strconv.Itoa(i+1))
	}
	// [1, 2, 3, 4, 5, 6]
	fmt.Println("Updated Slices", slice, "Size:", len(slice))

	slice2 := make([]string, len(slice))
	copy(slice2, slice)
	// [1, 2, 3, 4, 5, 6]
	fmt.Println("Copied Slices", slice2, "Size:", len(slice2))

	// [3, 4, 5] , 3
	fmt.Println("Index: 2-5 of Slice:", slice2[2:5], "Size:", len(slice2[2:5]))

	// [1, 2, 3, 4, 5], 5
	fmt.Println("Index: up to 5 of Slice:", slice[:5], "Size:", len(slice[:5]))
}
