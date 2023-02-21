package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func MakeHandler(instance *echo.Echo, s *resource) {
	instance.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Service is up.")
	})

	g := instance.Group("")
	baseUrl := "api/v1/product"

	g.POST(fmt.Sprintf("%s", baseUrl), s.CreateProduct)
}
