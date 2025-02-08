package database

import (
	"backend/internal/logger"
	"context"
	"time"

	"github.com/uptrace/bun"
)

type QueryHook struct{}

var log = logger.GetSugaredLogger()

func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	log.Infof("Execution Time %+v with query:\n %+v\n", time.Since(event.StartTime), string(event.Query))
}
