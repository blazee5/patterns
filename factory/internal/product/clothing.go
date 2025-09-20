package product

type ShirtFactory struct{}

func (sf ShirtFactory) CreateProduct() Product {
	return Shirt{}
}
