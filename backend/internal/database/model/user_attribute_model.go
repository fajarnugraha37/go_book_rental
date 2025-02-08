package model

import (
	"backend/internal/database/model/base"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ bun.BeforeInsertHook = (*UserAttribute)(nil)
var _ bun.BeforeUpdateHook = (*UserAttribute)(nil)

type UserAttribute struct {
	bun.BaseModel `bun:"table:user_attributes,alias:u"`

	ID     uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	UserID uuid.UUID `bun:"user_id,unique:user_attribute_key_per_user,type:uuid,notnull"`
	Key    string    `bun:"key,unique:user_attribute_key_per_user,notnull"`
	Type   string    `bun:"type,notnull"`
	Value  string    `bun:"value,notnull"`

	base.AuditColumn
	User *User `bun:"rel:belongs-to,join:user_id=id"`
}
