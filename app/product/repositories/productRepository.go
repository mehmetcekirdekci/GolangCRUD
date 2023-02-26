package repositories

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Create(product types.Product) error
		GetById(id int) (*types.Product, error)
		Update(product types.Product) error
		Delete(id int) error
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
	err := receiver.db.Table(Products).Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (receiver *productRepository) GetById(id int) (*types.Product, error) {
	var product types.Product
	err := receiver.db.Table(Products).Where(types.Product{Id: id}).Find(&product).Error
	if err != nil {
		return nil, err
	}
	if &product == nil || product.Id == 0 {
		return nil, nil
	}
	return &product, nil
}

func (receiver *productRepository) Update(product types.Product) error {
	err := receiver.db.Table(Products).Updates(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (receiver *productRepository) Delete(id int) error {
	err := receiver.db.Table(Products).Delete(&types.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
