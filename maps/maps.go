package main

import "fmt"

func main() {
	// map[<key-type>]<val-type>
	maps := make(map[string]int)

	maps["v1"] = 10
	maps["v2"] = 23
	maps["v3"] = 7
	// [v1:10, v2:23, v3:7]
	fmt.Println("MAP CONTENT:", maps)

	v1Copy := maps["v1"]
	fmt.Println("v1Copy VALUE:", v1Copy) // 10
	fmt.Println("LENGTH OF maps:", len(maps)) // 3

	delete(maps, "v2")
	// [v1:10, v3:7]
	fmt.Println("MAP CONTENT:", maps)
	fmt.Println("MAP LENGTH:", len(maps)) // 2

	// Second Return(bool): if key exists in map
	_, present := maps["v2"]
	fmt.Println("PRESENT:", present) // false
	_, present = maps["v1"]
	fmt.Println("PRESENT:", present) // true

	// Initialize with specified values
	maps2 := map[string]int{"GoLang":10, "Ruby":9, "Python":10, "JS": 8}
	fmt.Println("CONTENT OF MAP2:", maps2)
	fmt.Println("LENGTH OF MAP2:", len(maps2))
}
