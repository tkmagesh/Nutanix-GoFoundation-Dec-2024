package main

// HINT: use strconv.Atoi() to convert strings to ints

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10))                //=> 10
	fmt.Println(sum("10", 20))          //=> 30
	fmt.Println(sum(10, 20, "30", 40))  //=> 100
	fmt.Println(sum(10, 20, "abc", 40)) //=> 70
	fmt.Println(sum())                  //=> 0
}

func sum(list ...any) int {
	var result int
	for _, x := range list {
		switch val := x.(type) {
		case int:
			result += val
		case string:
			if no, err := strconv.Atoi(val); err == nil {
				result += no
			}
		}
	}
	return result
}
