package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log := logger.GetSugaredLogger()
	defer log.Sync()

	cfg := config.LoadConfig()
	db := database.GetBunDB(cfg)
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
		result := db.QueryRow("SELECT current_timestamp")
		var res any
		result.Scan(&res)
		return c.JSON(map[string]any{
			"message": "Hello, World!",
			"result":  res,
		})
	})

	app.All("/api", func(c *fiber.Ctx) error {
		return c.SendString("API Hello, World!")
	})

	log.Fatal(app.Listen(":" + string(cfg.App.Port)))
}
