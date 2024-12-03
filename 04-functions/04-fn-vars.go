package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()
	sayHi = func() {
		fmt.Println("Hi World!")
	}
	sayHi()

	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	sayHello("Magesh")

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	addResult := add(100, 200)
	fmt.Println(addResult)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient int, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
