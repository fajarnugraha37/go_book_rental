package database

import (
	"backend/internal/config"
	"backend/internal/database/model"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// Service represents a service that interacts with a database.
type Database interface {
	Health() map[string]string
	Close() error
	GetSql() *sql.DB
	GetBun() *bun.DB
}

type connection struct {
	sql *sql.DB
	bun *bun.DB
}

var conn *connection

func SingletonConnection(cfg *config.Config) Database {
	var (
		isSqlNil = conn == nil || conn.sql == nil || conn.sql.Ping() != nil
		isBunNil = conn == nil || conn.bun == nil || conn.bun.Ping() != nil
	)
	if isSqlNil || isBunNil {
		conn = NewConnection(cfg).(*connection)
	}

	return conn
}

func NewConnection(cfg *config.Config) Database {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
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

	bundb := bun.NewDB(
		sqldb,
		pgdialect.New(),
	)
	bundb.RegisterModel(
		(*model.UserToRole)(nil),
	)
	bundb.AddQueryHook(
		&QueryHook{},
	)
	if err := bundb.Ping(); err != nil {
		panic(err)
	}

	return &connection{
		sql: sqldb,
		bun: bundb,
	}
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *connection) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.sql.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.sql.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *connection) Close() error {
	// log.Printf("Disconnected from database")
	return s.sql.Close()
}

func (s *connection) GetSql() *sql.DB {
	return s.sql
}

func (s *connection) GetBun() *bun.DB {
	return s.bun
}
