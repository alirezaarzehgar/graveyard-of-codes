package controller

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/alirezaarzehgar/hi/model"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	pass := fmt.Sprintf("%x", sha256.Sum256([]byte(c.FormValue("pass"))))
	user := model.User{
		Username: c.FormValue("user"),
		Password: pass,
	}
	r := db.Where(&user).Find(&user)
	if r.Error != nil || r.RowsAffected == 0 {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, map[string]any{
		"user_id": user.ID,
	})
}

func Register(c echo.Context) error {
	pass := fmt.Sprintf("%x", sha256.Sum256([]byte(c.FormValue("pass"))))
	user := model.User{
		Username: c.FormValue("user"),
		Password: pass,
	}
	err := db.Create(&user).Error
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, map[string]any{
		"user_id": user.ID,
	})
}

func Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Not implemented",
	})
}
