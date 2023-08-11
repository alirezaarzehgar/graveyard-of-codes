package api

import (
	"encoding/json"
	"hi/model"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AddProduct(c echo.Context) error {
	var p model.Product
	if err := json.NewDecoder(c.Request().Body).Decode(&p); err != nil {
		return echo.ErrBadRequest
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.RegisteredClaims)

	id, _ := strconv.Atoi(claims.ID)
	p.OwnerID = uint(id)

	if err := model.CreateProduct(&p); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	p, err := model.GetProductByID(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, p)
}

func EditProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	var p model.Product
	if err := json.NewDecoder(c.Request().Body).Decode(&p); err != nil {
		return echo.ErrBadRequest
	}

	if err := model.EditProductByID(uint(id), &p); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	if err := model.DeleteProductByID(uint(id)); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
