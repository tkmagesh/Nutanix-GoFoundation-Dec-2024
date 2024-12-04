package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
		productRanks["pen"] = 3
		productRanks["pencil"] = 1
		productRanks["marker"] = 5
	*/

	// var productRanks map[string]int = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// var productRanks = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// productRanks := map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	productRanks := map[string]int{
		"pen":    3,
		"pencil": 1,
		"marker": 5,
	}
	fmt.Println(productRanks)

	fmt.Println("# of product ranks :", len(productRanks))

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// checking for the existence of a key
	// keyToCheck := "notepad" //non-existent key
	keyToCheck := "pen"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q = %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}

	// Removing an item
	// keyToRemove := "pen"
	keyToRemove := "notepad" //non-existent key
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing key %q\n", keyToRemove)
	fmt.Println("productRanks =", productRanks)
}
