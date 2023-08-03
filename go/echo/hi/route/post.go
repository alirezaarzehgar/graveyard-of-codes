package route

import (
	"github.com/alirezaarzehgar/hi/controller"
	"github.com/labstack/echo/v4"
)

func InitPost(r *echo.Echo) {
	g := r.Group("/post")
	g.POST("", controller.CreatePost)
	g.GET("/:id", controller.GetPost)
	g.PUT("/:id", controller.EditPost)
	g.DELETE("/:id", controller.DeletePost)
}
