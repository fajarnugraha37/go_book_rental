package server

import (
	"backend/internal/config"
	"backend/internal/database"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	*fiber.App
	cfg *config.Config
	db  database.Database
}

func New() *Server {
	var (
		cfg    = config.LoadConfig()
		server = &Server{
			App: fiber.New(fiber.Config{
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
			}),
			cfg: cfg,
			db:  database.SingletonConnection(cfg),
		}
	)

	return server
}

func (srv *Server) StartListening() {
	err := srv.Listen(fmt.Sprintf(":%v", srv.cfg.App.Port))
	if err != nil {
		panic(errors.New(fmt.Sprintf("http server error: %s", err)))
	}
}
