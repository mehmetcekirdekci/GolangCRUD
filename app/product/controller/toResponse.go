package controller

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
)

func ToBaseResponse(statusCode int, success bool, message string) *BaseResponse {
	baseResponse := new(BaseResponse)
	baseResponse.StatusCode = statusCode
	baseResponse.Success = success
	baseResponse.Message = message
	return baseResponse
}

func ToGetProductResponse(product *types.Product, statusCode int, success bool, message string) *GetProductResponse {
	getProductResponse := new(GetProductResponse)
	getProductResponse.BaseResponse.Success = success
	getProductResponse.BaseResponse.Message = message
	getProductResponse.BaseResponse.StatusCode = statusCode
	if product != nil {
		data := new(GetProductResponseDataItem)
		data.Product = *product
		getProductResponse.Data = data
	}
	return getProductResponse
}
