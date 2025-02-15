package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"context"
	"fmt"

	"github.com/uptrace/bun/migrate"
)

// A collection of migrations.
var migrations = migrate.NewMigrations()

func execute() {
	if err := migrations.DiscoverCaller(); err != nil {
		panic(err)
	}

	var (
		cfg      = config.LoadConfig()
		db       = database.SingletonConnection(cfg).GetBun()
		migrator = migrate.NewMigrator(db, migrations)
		ctx      = context.Background()
	)

	// create migration tables
	migrator.Init(ctx)

	// migrate database
	if err := migrator.Lock(ctx); err != nil {
		panic(err)
	}
	defer migrator.Unlock(ctx)

	group, err := migrator.Migrate(ctx)
	if err != nil {
		panic(err)
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
		return
	}

	fmt.Printf("migrated to %s\n", group)
	result, err := migrator.AppliedMigrations(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("result %+v", result)
}
