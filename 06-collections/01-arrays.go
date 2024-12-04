package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// Iterating using indexer
	fmt.Println("Iterating using indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// Iterating using range
	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	nos2 := nos //creating a copy of the "nos" array
	nos2[0] = 9999
	fmt.Printf("nos[0] = %d, nos2[0] = %d\n", nos[0], nos2[0])

	var nosPtr *[5]int
	nosPtr = &nos
	// fmt.Println((*nosPtr)[0])
	fmt.Println(nosPtr[0]) // accessing the members through a pointer of a collection does not require dereferencing

	fmt.Println("Before sorting, nos = ", nos)
	sort(&nos)
	fmt.Println("After sorting, nos = ", nos)
}

func sort(values *[5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println("[sort] values = ", values)
}
