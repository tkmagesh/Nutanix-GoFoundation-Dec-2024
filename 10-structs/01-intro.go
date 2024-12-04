package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func main() {
	/*
		var pen Product
		pen.Id = 100
		pen.Name = "Fountain Pen"
		pen.Cost = 20
	*/

	// var pen Product = Product{100, "Fountain Pen", 20} // advisable
	/*
		var pen Product = Product{
			Id:   100,
			Name: "Fountain Pen",
			Cost: 20,
		}
	*/
	pen := Product{
		Id:   100,
		Name: "Fountain Pen",
		Cost: 20,
	}

	// fmt.Printf("%#v\n", pen)
	fmt.Printf("%+v\n", pen)

	p2 := pen // creating a copy coz struct instances are just values
	p2.Cost = 100
	fmt.Printf("pen.Cost = %v, p2.Cost = %v\n", pen.Cost, p2.Cost)

	ApplyDiscount(&pen, 10)
	fmt.Println("After applying 10% discount, pen =", pen)

	pp1 := Product{Id: 200, Name: "pencil", Cost: 5}
	pp2 := Product{Id: 200, Name: "pencil", Cost: 5}
	fmt.Println(pp1 == pp2) //struct instances are compared by values
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
