package main

import (
	"backend/internal/logger"
	"backend/internal/server"
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var (
		srv = server.New()
		log = logger.GetSugaredLogger()
	)
	defer log.Sync()

	srv.All("/", func(c *fiber.Ctx) error {
		var res any
		return c.JSON(map[string]any{
			"message": "Hello, World!",
			"result":  res,
		})
	})

	srv.All("/api", func(c *fiber.Ctx) error {
		return c.SendString("API Hello, World!")
	})

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	go srv.StartListening()

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(srv.App, done)

	// Wait for the graceful shutdown to complete
	<-done
	log.Info("Graceful shutdown complete.")
}

func gracefulShutdown(fiberApp *fiber.App, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := fiberApp.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}
