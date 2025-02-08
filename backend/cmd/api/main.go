package main

import (
	"backend/internal/config"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()
	app := fiber.New(fiber.Config{
		ServerHeader:             cfg.App.Name,
		AppName:                  cfg.App.Name,
		CaseSensitive:            false,
		StrictRouting:            true,
		BodyLimit:                4 * 1024 * 1024,
		Concurrency:              256 * 1024,
		EnablePrintRoutes:        true,
		EnableSplittingOnParsers: true,
		WriteTimeout:             time.Second * 60,
		ReadTimeout:              time.Second * 60,
	})

	app.All("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.All("/api", func(c *fiber.Ctx) error {
		return c.SendString("API Hello, World!")
	})

	log.Fatal(app.Listen(":" + string(cfg.App.Port)))
}
