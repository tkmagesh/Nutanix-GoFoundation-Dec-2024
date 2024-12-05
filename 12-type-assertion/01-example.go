package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Quis eu voluptate ullamco amet anim labore ullamco aute anim magna cupidatat sit."
	x = true
	x = 99.99
	x = struct{}{}
	fmt.Println(x)

	// x = 100
	x = "In adipisicing amet occaecat enim dolore sint duis do deserunt nisi est."
	// y := x * 2

	// NOT SAFE
	/*
		y := x.(int) * 2
		fmt.Println(y)
	*/

	// SAFE
	x = "In adipisicing amet occaecat enim dolore sint duis do deserunt nisi est."
	// x = 200
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// using type-switch
	x = 100
	x = "Quis eu voluptate ullamco amet anim labore ullamco aute anim magna cupidatat sit."
	x = true
	x = 99.99
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a bool, !x = ", !val)
	case float64:
		fmt.Println("x is a float64, x * 0.99 =", val*0.99)
	default:
		fmt.Println("x is an unknown type!")
	}

}
