package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/database/model"
	"backend/internal/database/repo"
	"backend/internal/logger"
	"backend/pkg/helper"
	"context"
	"strconv"
	"time"
)

func main() {
	log := logger.GetSugaredLogger()
	defer log.Sync()

	cfg := config.LoadConfig()
	db := database.GetBunDB(cfg)
	ur := repo.NewUserRepo(db)

	for i := 0; i < 5; i++ {
		u := &model.User{
			Username: "username_" + time.Now().Format(time.StampMilli) + "_" + strconv.Itoa(i),
			IsActive: false,
		}
		ur.Insert(context.Background(), u)
	}

	user, err := ur.FindAll(
		context.Background(),
		repo.UserFilter{
			Username: helper.ToPtr("%779_2"),
			// AuditFilter: repo.AuditFilter{
			// 	// ID: helper.ToPtr("01a0c39a-5f77-4028-8081-7dd4f2d6fc07"),
			// },
		},
		nil,
		nil,
		// &repo.Pageable{
		// 	PageSize: 1,
		// },
		// &repo.Sortable{
		// 	SortBy:    "created_at",
		// 	Direction: 1,
		// },
	)
	log.Infof("user %+v", user)
	log.Infof("error %+v", err)

	// app := fiber.New(fiber.Config{
	// 	ServerHeader:             cfg.App.Name,
	// 	AppName:                  cfg.App.Name,
	// 	CaseSensitive:            false,
	// 	StrictRouting:            true,
	// 	BodyLimit:                4 * 1024 * 1024,
	// 	Concurrency:              256 * 1024,
	// 	EnablePrintRoutes:        true,
	// 	EnableSplittingOnParsers: true,
	// 	WriteTimeout:             time.Second * 60,
	// 	ReadTimeout:              time.Second * 60,
	// })

	// app.All("/", func(c *fiber.Ctx) error {
	// 	result := db.QueryRow("SELECT current_timestamp")
	// 	var res any
	// 	result.Scan(&res)
	// 	return c.JSON(map[string]any{
	// 		"message": "Hello, World!",
	// 		"result":  res,
	// 	})
	// })

	// app.All("/api", func(c *fiber.Ctx) error {
	// 	return c.SendString("API Hello, World!")
	// })

	// log.Fatal(app.Listen(":" + string(cfg.App.Port)))
}
