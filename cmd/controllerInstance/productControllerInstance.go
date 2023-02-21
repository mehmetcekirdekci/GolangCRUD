package controllerInstance

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/controller"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/services"
	"net/http"
)

func ProductControllerInstance(productBaseService services.BaseService, instance *echo.Echo) {
	getProductControllerInstance(productBaseService, instance)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func setCustomValidator(instance *echo.Echo) {
	instance.Validator = &CustomValidator{validator: validator.New()}
}

func getProductControllerInstance(productBaseService services.BaseService, instance *echo.Echo) {
	setCustomValidator(instance)
	controller.MakeHandler(instance, controller.NewController(productBaseService))
}
