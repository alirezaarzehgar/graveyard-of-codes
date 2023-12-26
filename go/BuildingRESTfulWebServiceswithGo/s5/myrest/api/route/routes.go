package route

import (
	"myrest/api/handler"
	_ "myrest/docs"

	"github.com/labstack/echo/v4"
	es "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/swagger/*", es.WrapHandler)

	e.POST("/user/new", handler.CreateUser)
	e.GET("/user/:name", handler.GetUser)
	e.GET("/user/all", handler.GetAllUsers)
	e.PUT("/user/:name", handler.UpdateUser)
	e.DELETE("/user/:name", handler.DeleteUser)

	return e
}
