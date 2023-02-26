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
	baseUrl := "api/v1/user"

	g.POST(fmt.Sprintf("%s", baseUrl), s.CreateUser)
	g.PUT(fmt.Sprintf("%s", baseUrl), s.UpdateUser)
	g.GET(fmt.Sprintf("%s", baseUrl), s.GetUser)
	g.DELETE(fmt.Sprintf("%s", baseUrl), s.DeleteUser)
}
