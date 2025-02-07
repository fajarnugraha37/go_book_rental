package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.All("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.All("/api", func(c *fiber.Ctx) error {
		return c.SendString("API Hello, World!")
	})

	log.Fatal(app.Listen(":8080"))
}
