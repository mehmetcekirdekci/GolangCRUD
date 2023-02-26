package controller

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
)

func (request CreateProductRequest) FromCreateProductRequestToCreateProductDto() *types.CreateProductDto {
	createProductDto := new(types.CreateProductDto)
	createProductDto.Name = request.Name
	createProductDto.Price = request.Price
	createProductDto.Currency = request.Currency
	return createProductDto
}

func (request UpdateProductRequest) FromUpdateProductRequestToUpdateProductDto() *types.UpdateProductDto {
	updateProductDto := new(types.UpdateProductDto)
	updateProductDto.Id = request.Id
	updateProductDto.Name = request.Name
	updateProductDto.Price = request.Price
	updateProductDto.Currency = request.Currency
	return updateProductDto
}
