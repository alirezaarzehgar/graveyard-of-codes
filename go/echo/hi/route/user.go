package route

import (
	"github.com/alirezaarzehgar/hi/controller"
	"github.com/labstack/echo/v4"
)

func InitUser(r *echo.Echo) {
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/logout", controller.Logout)
}
