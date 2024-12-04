package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred, err :", err)
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(x, y int) (quotient, remainder int) {
	fmt.Println("calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	fmt.Println("calculating remainder")
	remainder = x % y
	return
}
