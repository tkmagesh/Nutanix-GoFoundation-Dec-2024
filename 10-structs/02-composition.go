package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

type PerishableProduct struct {
	Product // PerishableProduct "IS A" Product
	Expiry  string
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	// var milk PerishableProduct = PerishableProduct{Product{200, "Milk", 20}, "2 Days"}
	var milk PerishableProduct = PerishableProduct{
		Product: Product{ // "Product" is both the composed type & attribute name
			Id:   200,
			Name: "Milk",
			Cost: 20,
		},
		Expiry: "2 Days",
	}

	// Accessing the attributes
	// fmt.Println(milk.Product.Id, milk.Product.Name, milk.Product.Cost, milk.Expiry)
	fmt.Println(milk.Id, milk.Name, milk.Cost, milk.Expiry)
	fmt.Printf("%+v\n", milk)

	// Using the "factory"
	eggs := NewPerishableProduct(201, "Eggs", 15, "2 Weeks")
	fmt.Println(eggs)

	ApplyDiscount(&eggs.Product, 10)
	fmt.Println("After applying 10% discount")
	fmt.Println(eggs)
}

func ApplyDiscount(p *Product, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
