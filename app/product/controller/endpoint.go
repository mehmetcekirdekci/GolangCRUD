package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/services"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
	"net/http"
	"strconv"
)

type resource struct {
	baseService services.BaseService
}

func NewController(baseService services.BaseService) *resource {
	return &resource{
		baseService: baseService,
	}
}

// CreateProduct godoc
// @Summary Creates product
// @Tags product
// @Param request body CreateProductRequest true "CreateProductRequest"
// @Success 201 {object} BaseResponse
// @Failure 400 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /api/v1/product [post]
func (receiver *resource) CreateProduct(c echo.Context) error {
	request := new(CreateProductRequest)
	response := new(BaseResponse)
	var resultMessage string
	err := c.Bind(request)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	err = c.Validate(request)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	createProductDto := request.FromCreateProductRequestToCreateProductDto()
	resultMessage, err = receiver.baseService.Create(*createProductDto)
	if err != nil {
		response = ToBaseResponse(http.StatusInternalServerError, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}
	response = ToBaseResponse(http.StatusCreated, true, resultMessage)
	return c.JSON(response.StatusCode, response)
}

// UpdateProduct godoc
// @Summary Updates product by Id
// @Tags product
// @Param request body UpdateProductRequest true "UpdateProductRequest"
// @Failure 201 {object} BaseResponse
// @Failure 400 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /api/v1/product [put]
func (receiver *resource) UpdateProduct(c echo.Context) error {
	request := new(UpdateProductRequest)
	response := new(BaseResponse)
	var resultMessage string
	err := c.Bind(request)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	err = c.Validate(request)
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	updateProductDto := request.FromUpdateProductRequestToUpdateProductDto()
	var updatedProduct *types.Product
	updatedProduct, resultMessage, err = receiver.baseService.Update(*updateProductDto)
	if err != nil && updatedProduct == nil || updatedProduct.Id == 0 {
		response = ToBaseResponse(http.StatusNotFound, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	if err != nil {
		response = ToBaseResponse(http.StatusInternalServerError, false, resultMessage)
		return echo.NewHTTPError(response.StatusCode, response)
	}

	response = ToBaseResponse(http.StatusCreated, true, resultMessage)
	return c.JSON(response.StatusCode, response)

}

// GetProduct godoc
// @Summary Gets product
// @Tags product
// @Param query int id true "id"
// @Success 201 {object} GetProductResponse
// @Failure 400 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /api/v1/product [get]
func (receiver *resource) GetProduct(c echo.Context) error {
	response := new(GetProductResponse)
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		response := ToGetProductResponse(nil, http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.BaseResponse.StatusCode, response)
	}

	product, resultMessage, err := receiver.baseService.GetById(id)
	if err != nil && product == nil || product.Id == 0 {
		response := ToGetProductResponse(product, http.StatusNotFound, false, err.Error())
		return echo.NewHTTPError(response.BaseResponse.StatusCode, response)
	}

	if err != nil {
		response := ToGetProductResponse(nil, http.StatusInternalServerError, false, resultMessage)
		return echo.NewHTTPError(response.BaseResponse.StatusCode, response)
	}

	response = ToGetProductResponse(product, http.StatusOK, true, resultMessage)
	return c.JSON(http.StatusOK, response)
}

// DeleteProduct godoc
// @Summary Deletes product
// @Tags product
// @Param query int id true "id"
// @Success 200 {object} BaseResponse
// @Failure 400 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /api/v1/product [delete]
func (receiver *resource) DeleteProduct(c echo.Context) error {
	response := new(BaseResponse)
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		response = ToBaseResponse(http.StatusBadRequest, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	product, resultMessage, err := receiver.baseService.Delete(id)
	if err != nil && product == nil || product.Id == 0 {
		response = ToBaseResponse(http.StatusNotFound, false, err.Error())
		return echo.NewHTTPError(response.StatusCode, response)
	}

	if err != nil {
		response := ToBaseResponse(http.StatusInternalServerError, false, resultMessage)
		return echo.NewHTTPError(response.StatusCode, response)
	}

	response = ToBaseResponse(http.StatusOK, true, resultMessage)
	return c.JSON(http.StatusOK, response)
}
