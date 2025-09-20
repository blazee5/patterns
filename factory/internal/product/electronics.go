package product

type SmartphoneFactory struct{}

func (sf SmartphoneFactory) CreateProduct() Product {
	return Smartphone{}
}

type LaptopFactory struct{}

func (lf LaptopFactory) CreateProduct() Product {
	return Laptop{}
}
