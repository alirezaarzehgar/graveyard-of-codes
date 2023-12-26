package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "hi/docs"
)

//	@version	2.0
//	@title		swagger test

func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":3000"))
}

type HTTPError struct {
	Status  int
	Error   string
	Message string
}

// Handler godoc
//
//	@Router			/handler [get]
//	@Summary		Just for fun
//	@Description	yooyoy
//	@Tags			fun
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	int
//	@Failure		400	{object}	HTTPError
//	@Failure		401	{object}	HTTPError
//	@Failure		402	{object}	HTTPError
//	@Failure		500	{object}	HTTPError
func Handler(c *echo.Context) error {

	return nil
}
