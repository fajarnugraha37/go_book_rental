package database

import (
	"backend/internal/config"
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var sqldb *sql.DB
var bundb *bun.DB

func connect(cfg config.Config) {
	sqldb = sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(fmt.Sprintf("%s:%d", cfg.Database.Host, cfg.Database.Port)),
		pgdriver.WithUser(cfg.Database.User),
		pgdriver.WithPassword(cfg.Database.Password),
		pgdriver.WithDatabase(cfg.Database.Name),
		pgdriver.WithApplicationName(cfg.App.Name),
		pgdriver.WithInsecure(true),
		pgdriver.WithTimeout(60*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(60*time.Second),
		pgdriver.WithWriteTimeout(60*time.Second),
	))
	if err := sqldb.Ping(); err != nil {
		panic(err)
	}

	bundb = bun.NewDB(sqldb, pgdialect.New())
	if err := bundb.Ping(); err != nil {
		panic(err)
	}
}

func GetSqlDB(cfg config.Config) *sql.DB {
	if sqldb == nil || sqldb.Ping() != nil {
		connect(cfg)
	}

	return sqldb
}

func GetBunDB(cfg config.Config) *bun.DB {
	if bundb == nil || bundb.Ping() != nil {
		connect(cfg)
	}

	return bundb
}
