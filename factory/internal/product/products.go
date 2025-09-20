package product

import "fmt"

type Smartphone struct{}

func (s Smartphone) Order() string {
	return fmt.Sprintf("Ordering smartphone: iPhone 15 - $999.99")
}

type Laptop struct{}

func (l Laptop) Order() string {
	return fmt.Sprintf("Ordering laptop: MacBook Pro - $2499.99")
}

type Shirt struct{}

func (s Shirt) Order() string {
	return fmt.Sprintf("Ordering shirt: Cotton Polo - $49.99")
}
