package main

import (
	"backend/internal/database/model"
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func init() {
	cols := []string{"is_active", "created_at", "updated_at", "deleted_flag"}
	migrations.Add(migrate.Migration{
		Name:    "00000000000002_create_users_table",
		Comment: "user table creations",
		Up: func(ctx context.Context, db *bun.DB) error {
			_, err := db.NewCreateTable().
				Model((*model.User)(nil)).
				IfNotExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			for _, col := range cols {
				_, err := db.NewCreateIndex().
					Model((*model.User)(nil)).
					IfNotExists().
					Index(col + "__users_idx").
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
				Model((*model.User)(nil)).
				IfExists().
				Cascade().
				Restrict().
				Exec(ctx)

			return err
		},
	})
}

// WithForeignKeys().
// ForeignKey(`(fkey) REFERENCES table1 (pkey) ON DELETE CASCADE`).
// PartitionBy("HASH (id)").
// TableSpace("fasttablespace").
