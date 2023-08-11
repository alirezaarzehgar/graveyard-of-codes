package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hi(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK!")
}
