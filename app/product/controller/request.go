package controller

import "github.com/mehmetcekirdekci/GolangCRUD/app/product/types"

type (
	CreateProductRequest struct {
		Name     string                 `json:"name" validate:"required"`
		Price    float32                `json:"price" validate:"required,gt=0"`
		Currency types.CurrencyTypeEnum `json:"currency" validate:"required,gte=0,lte=2"`
	}
)
