package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

func main() {

	pen := Product{
		Id:   100,
		Name: "Fountain Pen",
		Cost: 20,
	}

	/*
		fmt.Println(Format(pen))
		ApplyDiscount(&pen, 10)
		fmt.Println("After applying 10% discount")
		fmt.Println(Format(pen))
	*/

	fmt.Println(pen.Format())
	pen.ApplyDiscount(10)
	fmt.Println("After applying 10% discount")
	fmt.Println(pen.Format())

}

func (p Product) Format() string {
	return fmt.Sprintf("Id : %d, Name : %q, Cost : %0.02f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
