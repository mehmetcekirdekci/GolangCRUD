package repositories

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Create(product types.Product) error
	}

	productRepository struct {
		db *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (receiver *productRepository) Create(product types.Product) error {
	err := receiver.db.Table(Products).Create(product).Error
	return err
}
