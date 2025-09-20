package main

import (
	"fmt"

	"github.com/blazee5/patterns/internal/product"
)

func main() {
	// client code
	var f product.ProductFactory

	// smartphone part
	f = product.SmartphoneFactory{}
	smartphone := f.CreateProduct()
	fmt.Println(smartphone.Order())

	// laptop part
	f = product.LaptopFactory{}
	laptop := f.CreateProduct()
	fmt.Println(laptop.Order())

	// shirt part
	f = product.ShirtFactory{}
	shirt := f.CreateProduct()
	fmt.Println(shirt.Order())
}
