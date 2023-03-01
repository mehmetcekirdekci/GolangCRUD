package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
	mock_repositories "github.com/mehmetcekirdekci/GolangCRUD/mocks/product/repositories"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var mockProductRepository *mock_repositories.MockProductRepository
var testService BaseService

var mockProductList []types.Product

func setup(t *testing.T) func() {
	preLoadMockDatas()

	ct := gomock.NewController(t)
	defer ct.Finish()

	mockProductRepository = mock_repositories.NewMockProductRepository(ct)
	testService = NewBaseService(mockProductRepository)
	return func() {
		testService = nil
		defer ct.Finish()
	}
}

func preLoadMockDatas() {
	preLoadMockProducts()
}

func preLoadMockProducts() {
	for i := 1; i <= 5; i++ {
		mockProduct := new(types.Product)
		mockProduct.Id = i
		mockProduct.Name = fmt.Sprintf("Product%d", i)
		mockProduct.Price = float32(i * 100)
		mockProduct.Currency = types.CurrencyTypeEnum(rand.Intn(3))
		mockProductList = append(mockProductList, *mockProduct)
	}
}

func TestBaseService_Create(t *testing.T) {
	td := setup(t)
	defer td()

	mockProductDtoToBeCreated := types.CreateProductDto{
		Name:     fmt.Sprintf("Product%d", len(mockProductList)+1),
		Price:    float32((len(mockProductList) + 1) * 100),
		Currency: types.CurrencyTypeEnum(rand.Intn(3)),
	}

	mockProductRepository.EXPECT().Create(*mockProductDtoToBeCreated.FromCreateProductDtoToProduct()).Return(nil)
	resultMessage, err := testService.Create(mockProductDtoToBeCreated)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, resultMessage)
	assert.Equal(t, types.SuccesfulOperation, resultMessage)
	assert.Nil(t, err)
}

func TestBaseService_GetById(t *testing.T) {
	td := setup(t)
	defer td()

	id := rand.Intn(len(mockProductList))
	mockProductRepository.EXPECT().GetById(id).Return(&mockProductList[id-1], nil)

	product, resultMessage, err := testService.GetById(id)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, resultMessage)
	assert.NotEmpty(t, product)
	assert.Equal(t, types.SuccesfulOperation, resultMessage)
}
