package controllerInstance

import (
	"github.com/labstack/echo/v4"
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/controller"
	"github.com/mehmetcekirdekci/GolangCRUD/app/user/services"
)

func UserControllerInstance(userBaseService services.BaseService, instance *echo.Echo) {
	getUserControllerInstance(userBaseService, instance)
}

func getUserControllerInstance(userBaseService services.BaseService, instance *echo.Echo) {
	setCustomValidator(instance)
	controller.MakeHandler(instance, controller.NewController(userBaseService))
}
