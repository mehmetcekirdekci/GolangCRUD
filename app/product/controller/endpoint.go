package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/services"
	"net/http"
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
