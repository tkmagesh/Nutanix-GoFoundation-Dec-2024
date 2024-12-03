package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")
	fmt.Println(add(100, 200))
	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	/*
		q, _ := divide(100, 7)
		fmt.Printf("dividing 100 by 7, quotient = %d \n", q)
	*/
}

func sayHi() {
	fmt.Println("Hi there!")
}

func sayHello(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

func add(x, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// named results
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
