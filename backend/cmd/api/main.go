package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/database/repo"
	"backend/internal/database/repo/filter"
	"backend/internal/logger"
	"context"
)

func main() {
	log := logger.GetSugaredLogger()
	defer log.Sync()

	cfg := config.LoadConfig()
	db := database.GetBunDB(cfg)
	userRepo := repo.NewUserRepo(db)
	// userAttrRepo := repo.NewUserAttributeRepo(db)

	// for i := 0; i < 15; i++ {
	// 	u := &model.User{
	// 		Username: "uname_" + strconv.Itoa(i) + time.Now().GoString(),
	// 		IsActive: true,
	// 	}
	// 	userRepo.Insert(context.Background(), u)

	// 	userAttrRepo.Insert(context.Background(), &model.UserAttribute{
	// 		UserID: u.ID,
	// 		Key:    "first_name",
	// 		Type:   "string",
	// 		Value:  "first_name_" + strconv.Itoa(i),
	// 	})
	// 	userAttrRepo.Insert(context.Background(), &model.UserAttribute{
	// 		UserID: u.ID,
	// 		Key:    "last_name",
	// 		Type:   "string",
	// 		Value:  "last_name_" + strconv.Itoa(i),
	// 	})
	// 	userAttrRepo.Insert(context.Background(), &model.UserAttribute{
	// 		UserID: u.ID,
	// 		Key:    "business_name",
	// 		Type:   "string",
	// 		Value:  "business_name_" + strconv.Itoa(i),
	// 	})
	// }

	predicate := &filter.Predicate{
		Operation: filter.OperationEnum("or"),
		Filters: []filter.QueryFilter{
			{
				Field:      "users.is_active",
				Comparator: filter.EQ,
				Param:      true,
			},
		},
		Pageable: &filter.Pageable{
			Page:     1,
			PageSize: 1,
		},
		Sortable: &[]filter.Sortable{
			{
				SortBy:    "users.created_at",
				Direction: filter.DESC,
			},
		},
		Relations: []string{"UserAttributes"},
	}
	page, err := userRepo.FindPageable(context.Background(), predicate)
	log.Infof("Page %+v", page.Page)
	log.Infof("TotalItems %+v", page.TotalItems)
	log.Infof("page.TotalPages %+v", page.TotalPages)
	log.Infof("Items %+v", page.Items[0].UserAttributes[0].Key)
	log.Infof("Items %+v", page.Items[0].UserAttributes[0].Value)
	log.Infof("error %+v\n\n", err)
	// for {
	// 	page, err := userRepo.FindPageable(context.Background(), predicate)
	// 	log.Infof("Page %+v", page.Page)
	// 	log.Infof("len(page.Items) %+v", len(page.Items))
	// 	log.Infof("TotalItems %+v", page.TotalItems)
	// 	log.Infof("page.TotalPages %+v", page.TotalPages)
	// 	log.Infof("error %+v\n\n", err)
	// 	if err == nil && page.HasNext() {
	// 		predicate.Pageable.Page = page.NextPage
	// 	} else {
	// 		break
	// 	}
	// }

	// user, err := ur.FindAll(
	// 	context.Background(),
	// 	repo.UserFilter{
	// 		// Username: helper.ToPtr("%779_2"),
	// 		CommonFilter: repo.CommonFilter{
	// 			ID:        helper.ToPtr("0d81d764-23e0-4b8e-a08a-21df4f62ebda"),
	// 			Relations: []string{"UserAttributes"},
	// 		},
	// 	},
	// 	nil,
	// 	nil,
	// 	// &repo.Pageable{
	// 	// 	PageSize: 1,
	// 	// },
	// 	// &repo.Sortable{
	// 	// 	SortBy:    "created_at",
	// 	// 	Direction: 1,
	// 	// },
	// )
	// log.Infof("user %+v", user)
	// log.Infof("error %+v", err)

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
