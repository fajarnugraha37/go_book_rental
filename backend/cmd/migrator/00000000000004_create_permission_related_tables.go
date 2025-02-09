package main

import (
	"backend/internal/database/model"
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func init() {
	cols := []string{"created_at", "updated_at", "deleted_flag"}
	migrations.Add(migrate.Migration{
		Name:    "00000000000004_create_permission_related_tables",
		Comment: "user table creations",
		Up: func(ctx context.Context, db *bun.DB) error {
			// Role
			_, err := db.NewCreateTable().
				Model((*model.Role)(nil)).
				IfNotExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			for _, col := range cols {
				_, err := db.NewCreateIndex().
					Model((*model.Role)(nil)).
					IfNotExists().
					Index(col + "__role_idx").
					Column(col).
					Exec(ctx)
				if err != nil {
					return err
				}
			}

			// Permission
			_, err = db.NewCreateTable().
				Model((*model.Permission)(nil)).
				IfNotExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			return nil
		},
		Down: func(ctx context.Context, db *bun.DB) error {
			_, err := db.NewDropTable().
				Model((*model.Permission)(nil)).
				IfExists().
				Cascade().
				Restrict().
				Exec(ctx)
			if err != nil {
				return err
			}

			_, err = db.NewDropTable().
				Model((*model.Role)(nil)).
				IfExists().
				Cascade().
				Restrict().
				Exec(ctx)
			if err != nil {
				return err
			}

			return err
		},
	})
}
