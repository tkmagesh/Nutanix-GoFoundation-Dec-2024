package main

import "fmt"

func main() {
	// var nos []int
	nos := make([]int, 0 /* # of values to be initialized */, 3 /* total memory to be allocated */)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 3)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 1)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 4)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 2)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 5)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
}
