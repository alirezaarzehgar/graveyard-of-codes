package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
