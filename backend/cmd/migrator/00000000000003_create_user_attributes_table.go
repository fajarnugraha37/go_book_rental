package main

import (
	"backend/internal/database/model"
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func init() {
	cols := []string{"user_id", "key", "created_at", "updated_at", "deleted_flag"}
	migrations.Add(migrate.Migration{
		Name:    "00000000000003_create_user_attributes_table",
		Comment: "user table creations",
		Up: func(ctx context.Context, db *bun.DB) error {
			_, err := db.NewCreateTable().
				Model((*model.UserAttribute)(nil)).
				IfNotExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			for _, col := range cols {
				_, err := db.NewCreateIndex().
					Model((*model.UserAttribute)(nil)).
					IfNotExists().
					Index(col + "__user_attribute_idx").
					Column(col).
					Exec(ctx)
				if err != nil {
					return err
				}
			}

			return nil
		},
		Down: func(ctx context.Context, db *bun.DB) error {
			_, err := db.NewDropTable().
				Model((*model.UserAttribute)(nil)).
				IfExists().
				Cascade().
				Restrict().
				Exec(ctx)

			return err
		},
	})
}
