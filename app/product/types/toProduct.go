package types

func (createProductDto CreateProductDto) FromCreateProductDtoToProduct() *Product {
	product := new(Product)
	product.Name = createProductDto.Name
	product.Price = createProductDto.Price
	product.Currency = createProductDto.Currency
	return product
}
