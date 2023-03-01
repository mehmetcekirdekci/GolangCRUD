package services

import (
	"errors"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/repositories"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
)

type (
	BaseService interface {
		Create(createProductDto types.CreateProductDto) (string, error)
		Update(updateProductDto types.UpdateProductDto) (*types.Product, string, error)
		GetById(id int) (*types.Product, string, error)
		Delete(id int) (*types.Product, string, error)
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
	err := receiver.productRepository.Create(*product)
	if err == nil {
		isOperationSuccesful = true
	}
	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return resultMessage, err
}

func (receiver *baseService) Update(updateProductDto types.UpdateProductDto) (*types.Product, string, error) {
	var resultMessage string
	var isOperationSuccesful bool
	productToBeUpdated, err := receiver.productRepository.GetById(updateProductDto.Id)

	ProductNotFound := err == nil && productToBeUpdated == nil
	ProductFoundedSuccesfuly := err == nil && productToBeUpdated != nil && productToBeUpdated.Id != 0

	if ProductNotFound {
		err = errors.New(types.ObjectWasNotFound)
	} else if ProductFoundedSuccesfuly {
		isOperationSuccesful = true
	}

	if !isOperationSuccesful {
		resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
		return nil, resultMessage, err
	}

	updatedProduct := updateProductDto.FromUpdateProductDtoToProduct()
	err = receiver.productRepository.Update(*updatedProduct)
	if err != nil {
		isOperationSuccesful = false
	}

	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return updatedProduct, resultMessage, err
}

func (receiver *baseService) GetById(id int) (*types.Product, string, error) {
	var resultMessage string
	var isOperationSuccesful bool
	product, err := receiver.productRepository.GetById(id)

	ProductNotFound := err == nil && product == nil
	ProductFoundedSuccesfuly := err == nil && product != nil && product.Id != 0

	if ProductNotFound {
		err = errors.New(types.ObjectWasNotFound)
	} else if ProductFoundedSuccesfuly {
		isOperationSuccesful = true
	}

	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return product, resultMessage, err
}

func (receiver *baseService) Delete(id int) (*types.Product, string, error) {
	var resultMessage string
	var isOperationSuccesful bool
	productToBeDeleted, err := receiver.productRepository.GetById(id)

	ProductNotFound := err == nil && productToBeDeleted == nil
	ProductFoundedSuccesfuly := err == nil && productToBeDeleted != nil && productToBeDeleted.Id != 0

	if ProductNotFound {
		err = errors.New(types.ObjectWasNotFound)
	} else if ProductFoundedSuccesfuly {
		isOperationSuccesful = true
	}

	if !isOperationSuccesful {
		resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
		return nil, resultMessage, err
	}

	err = receiver.productRepository.Delete(id)
	if err != nil {
		isOperationSuccesful = false
	}

	resultMessage = ResultMessageBuilder(isOperationSuccesful, err)
	return productToBeDeleted, resultMessage, err
}

func getById(receiver *baseService, id int) (*types.Product, error) {
	product, err := receiver.productRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return product, err
}
