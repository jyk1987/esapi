package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jyk1987/es/log"
)

func main() {
	log.Log.Debug("test")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":8080")
}
