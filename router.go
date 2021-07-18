package testecho

import (
	"fmt"
	"testecho/controller"

	"github.com/labstack/echo/v4"
)

func NewApi(port int) {
	e := echo.New()
	e.GET("/", controller.Hello)
	e.GET("/getuser", controller.GetUsers)
	e.Start(fmt.Sprintf(":%v", port))
}
