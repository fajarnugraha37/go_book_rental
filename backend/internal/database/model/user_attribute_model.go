package model

import (
	"backend/internal/database/model/base"
	"backend/pkg/helper"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ bun.BeforeAppendModelHook = (*UserAttribute)(nil)

type UserAttribute struct {
	bun.BaseModel `bun:"table:user_attributes,alias:user_attributes"`

	ID     uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	UserID uuid.UUID `bun:"user_id,unique:user_attribute_key_per_user,type:uuid,notnull"`
	Key    string    `bun:"key,unique:user_attribute_key_per_user,notnull"`
	Type   string    `bun:"type,notnull"`
	Value  string    `bun:"value,notnull"`
	User   User      `bun:"rel:belongs-to,join:user_id=id"`

	base.AuditColumn
}

func (m *UserAttribute) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = helper.ToPtr(time.Now())
		if m.CreatedBy == nil {
			m.CreatedBy = helper.ToPtr("system@intranet")
		}
	case *bun.UpdateQuery:
		m.UpdatedAt = helper.ToPtr(time.Now())
		if m.UpdatedBy == nil {
			m.UpdatedBy = helper.ToPtr("system@intranet")
		}
	}
	return nil
}
