package controllerInstance

import (
	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/controller"
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/services"
)

func ProductControllerInstance(productBaseService services.BaseService, instance *echo.Echo) {
	getProductControllerInstance(productBaseService, instance)
}

func getProductControllerInstance(productBaseService services.BaseService, instance *echo.Echo) {
	setCustomValidator(instance)
	controller.MakeHandler(instance, controller.NewController(productBaseService))
}
