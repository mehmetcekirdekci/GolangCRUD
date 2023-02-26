package controller

import "github.com/mehmetcekirdekci/GolangCRUD/app/product/types"

type (
	BaseResponse struct {
		StatusCode int
		Success    bool
		Message    string
	}

	GetProductResponse struct {
		BaseResponse BaseResponse
		Data         *GetProductResponseDataItem
	}
	GetProductResponseDataItem struct {
		Product types.Product
	}
)
