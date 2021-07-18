package controller

import (
	"net/http"
	"testecho/service"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	us := service.UserService{}
	return c.JSON(http.StatusOK, us.GetUsers())
}

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world")
}
