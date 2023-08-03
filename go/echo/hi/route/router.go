package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "ok",
		"message": c.Path(),
		"action":  c.Request().Method,
	})
}

func Init() *echo.Echo {
	router := echo.New()

	InitUser(router)
	InitPost(router)

	return router
}
