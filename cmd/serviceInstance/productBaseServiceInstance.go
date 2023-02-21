package serviceInstance

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/repositories"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/services"
)

func GetProductBaseServiceInstance(productRepository repositories.ProductRepository) services.BaseService {
	return getProductBaseServiceInstance(productRepository)
}

func getProductBaseServiceInstance(productRepository repositories.ProductRepository) services.BaseService {
	return services.NewBaseService(productRepository)
}
