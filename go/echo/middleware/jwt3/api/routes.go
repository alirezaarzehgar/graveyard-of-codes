package api

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var JwtKey []byte

func InitRoutes() *echo.Echo {
	JwtKey = []byte(os.Getenv("JWT_KEY"))

	e := echo.New()
	e.POST("/register", Register)
	e.POST("/login", Login)

	conf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SigningKey: JwtKey,
	}
	g := e.Group("/", echojwt.WithConfig(conf))
	g.POST("refresh", Refresh)
	g.POST("product", AddProduct)
	g.GET("product/:id", GetProduct)
	g.PUT("product/:id", EditProduct)
	g.DELETE("product/:id", DeleteProduct)
	return e
}
