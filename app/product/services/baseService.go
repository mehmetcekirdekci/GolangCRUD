package services

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/repositories"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
)

type (
	BaseService interface {
		Create(createProductDto types.CreateProductDto) (string, error)
	}

	baseService struct {
		productRepository repositories.ProductRepository
	}
)

func NewBaseService(productRepository repositories.ProductRepository) BaseService {
	if productRepository == nil {
		return nil
	}
	return &baseService{
		productRepository: productRepository,
	}
}

func (receiver *baseService) Create(createProductDto types.CreateProductDto) (string, error) {
	var resultMessage string
	var isOperationSuccesful bool
	product := createProductDto.FromCreateProductDtoToProduct()
	err := create(receiver, *product)
	if err == nil {
		isOperationSuccesful = true
	}
	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return resultMessage, err
}

func create(receiver *baseService, product types.Product) error {
	err := receiver.productRepository.Create(product)
	if err != nil {
		return err
	}
	return nil
}
