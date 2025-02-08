package base

import (
	"time"
)

type AuditColumn struct {
	CreatedAt   *time.Time `bun:"created_at,type:timestamptz,default:current_timestamp"`
	CreatedBy   *string    `bun:"created_by,default:'system'"`
	UpdatedAt   *time.Time `bun:"updated_at,type:timestamptz"`
	UpdatedBy   *string    `bun:"updated_by"`
	DeletedFlag bool       `bun:"deleted_flag,default:false"`
	DeletedAt   *time.Time `bun:"deleted_at,type:timestamptz"`
	DeletedBy   *string    `bun:"deleted_by"`
}
