package main

import "fmt"

func main() {
	var x int
	x = 100

	var xPtr *int
	xPtr = &x //value -> address
	fmt.Println("&x = ", xPtr)

	// dereferencing (address -> value)
	fmt.Println(*xPtr)

	fmt.Println(x == *(&x))

	fmt.Println("Before incrementing, x = ", x)
	increment(&x)
	fmt.Println("After incrementing, x = ", x)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(val *int) {
	fmt.Println("&val =", val)
	*val++
}

func swap(x, y *int) { // no return result
	*x, *y = *y, *x
}
