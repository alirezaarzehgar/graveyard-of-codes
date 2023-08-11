package api

import (
	"encoding/json"
	"fmt"
	"hi/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func createToken(id string, user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ID:        id,
		Issuer:    user,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})

	strToken, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return strToken, nil
}

func Register(c echo.Context) error {
	var u model.User
	if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		return echo.ErrBadRequest
	}

	if err := model.RegisterUser(&u); err != nil {
		return err
	}

	strToken, err := createToken(fmt.Sprint(u.ID), u.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{"bearer": strToken})
}

func Login(c echo.Context) error {
	var u model.User
	if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		return echo.ErrBadRequest
	}

	if err := model.LoginUser(&u); err != nil {
		return err
	}

	strToken, err := createToken(fmt.Sprint(u.ID), u.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{"bearer": strToken})
}

func Refresh(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.RegisteredClaims)

	strToken, err := createToken(claims.ID, claims.Issuer)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{"bearer": strToken})
}
