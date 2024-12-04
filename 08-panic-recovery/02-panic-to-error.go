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
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r, err := divideAdapter(100, divisor)
		if err == ErrDivideByZero {
			fmt.Println("Do not attempt to divide by 0. Try again!")
			continue
		}
		if err != nil {
			fmt.Println("Unknown error :", err)
			break
		}
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		break
	}
}

// adapter to conver the panic into an error
func divideAdapter(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party
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
