package main

import "fmt"

func main() {
	/*
		var nos []int
		nos = append(nos, 3)
		nos = append(nos, 1, 4, 2, 5)
	*/

	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// appending another slice
	tens := []int{30, 10, 50}
	nos = append(nos, tens...)

	// Iterating using indexer
	fmt.Println("Iterating using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// Iterating using range
	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	fmt.Println("Slicing")
	fmt.Printf("nos[3:6] = %v\n", nos[3:6])
	fmt.Printf("nos[3:] = %v\n", nos[3:])
	fmt.Printf("nos[:6] = %v\n", nos[:6])

	nos2 := nos //creating a copy of the "nos" slice, but both nos & nos2 point to the same underlying array
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	fmt.Println("Before sorting, nos = ", nos)
	sort(nos)
	fmt.Println("After sorting, nos = ", nos)
}

func sort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println("[sort] values = ", values)
}
