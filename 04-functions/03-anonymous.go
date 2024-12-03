package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	addResult := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(addResult)

}
