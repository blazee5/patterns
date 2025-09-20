package factory

import "github.com/blazee5/patterns/internal/product"

type ProductFactory interface {
	CreateProduct() product.Product
}
