package base

import (
	"backend/pkg/helper"
	"context"
	"time"

	"github.com/uptrace/bun"
)

type AuditColumn struct {
	CreatedAt   *time.Time `bun:"created_at,default:current_timestamp"`
	CreatedBy   *string    `bun:"created_by,default:'system'"`
	UpdatedAt   *time.Time `bun:"updated_at"`
	UpdatedBy   *string    `bun:"updated_by"`
	DeletedFlag bool       `bun:"deleted_flag,default:false"`
	DeletedAt   *time.Time `bun:"deleted_at"`
	DeletedBy   *string    `bun:"deleted_by"`
}

// BeforeUpdate implements bun.BeforeUpdateHook.
func (m *AuditColumn) BeforeUpdate(ctx context.Context, query *bun.UpdateQuery) error {
	m.UpdatedAt = helper.ToPtr(time.Now())
	if m.UpdatedBy == nil {
		m.UpdatedBy = helper.ToPtr("system@intranet")
	}

	return nil
}

// BeforeInsert implements bun.BeforeInsertHook.
func (m *AuditColumn) BeforeInsert(ctx context.Context, query *bun.InsertQuery) error {
	m.CreatedAt = helper.ToPtr(time.Now())
	if m.CreatedBy == nil {
		m.CreatedBy = helper.ToPtr("system@intranet")
	}

	return nil
}
