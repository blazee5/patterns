package product

type Product interface {
	Order() string
}

type ProductFactory interface {
	CreateProduct() Product
}
