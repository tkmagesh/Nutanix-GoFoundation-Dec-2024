package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

// Methods of "Product"
func (p Product) Format() string {
	return fmt.Sprintf("Id : %d, Name : %q, Cost : %0.02f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
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

// overriding the Product.Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {

	// Using the "factory"
	eggs := NewPerishableProduct(201, "Eggs", 15, "2 Weeks")
	fmt.Println(eggs.Format())

	// ApplyDiscount() is inherited from "Product" type
	eggs.ApplyDiscount(10)
	fmt.Println("After applying 10% discount")
	fmt.Println(eggs.Format())
}
