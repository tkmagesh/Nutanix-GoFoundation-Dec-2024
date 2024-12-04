package main

import (
	"fmt"

	// "github.com/tkmagesh/nutanix-gofoundation-dec-2024/09-modularity-app/calculator"
	"github.com/fatih/color"
	calc "github.com/tkmagesh/nutanix-gofoundation-dec-2024/09-modularity-app/calculator"
	"github.com/tkmagesh/nutanix-gofoundation-dec-2024/09-modularity-app/calculator/utils"
)

func main() {
	fmt.Println("Modularity app executed")
	color.Yellow("Modularity app executed")

	greet("Magesh") // app.go file
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("Op Count :", calculator.OpCount())
	*/

	// using package alias
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Op Count :", calc.OpCount())

	// using the nested package
	fmt.Println("Is 21 even ?:", utils.IsEven(21))
}
