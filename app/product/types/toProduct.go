package types

func (createProductDto CreateProductDto) FromCreateProductDtoToProduct() *Product {
	product := new(Product)
	//product.Name = createProductDto.Name
	product.Price = createProductDto.Price
	product.Currency = createProductDto.Currency
	return product
}

func (updateProductDto UpdateProductDto) FromUpdateProductDtoToProduct() *Product {
	product := new(Product)
	product.Id = updateProductDto.Id
	product.Name = updateProductDto.Name
	product.Price = updateProductDto.Price
	product.Currency = updateProductDto.Currency
	return product
}
