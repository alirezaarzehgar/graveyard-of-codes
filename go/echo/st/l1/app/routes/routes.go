package routes

import (
	"example/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/hi", controllers.Hi)
	return e
}
