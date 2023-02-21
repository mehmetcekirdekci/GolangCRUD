package repositoryInstance

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/repositories"
	"gorm.io/gorm"
)

func GetProductRepositoryInstance(postgre *gorm.DB) repositories.ProductRepository {
	return getProductRepositoryInstance(postgre)
}

func getProductRepositoryInstance(postgre *gorm.DB) repositories.ProductRepository {
	return repositories.NewProductRepository(postgre)
}
