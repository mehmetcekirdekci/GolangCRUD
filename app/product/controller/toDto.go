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
