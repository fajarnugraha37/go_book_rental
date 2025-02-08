package main

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func init() {
	migrations.Add(migrate.Migration{
		Name:    "00000000000001_uuid_extensions",
		Comment: "add uuid extension",
		Up: func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
			return err
		},
		Down: func(ctx context.Context, db *bun.DB) error {
			return nil
		},
	})
}
