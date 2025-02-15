package model

import (
	"backend/internal/database/model/base"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Permission struct {
	bun.BaseModel `bun:"table:premissions,alias:premissions"`

	ID         uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	RoleID     string            `bun:"role_id"`
	Resource   string            `bun:"resource,notnull"`
	Privilege  map[string]string `bun:"privilege,notnull,type:jsonb,default:'{}'"`
	Expression *string           `bun:"Expression"`
	Role       Role              `bun:"rel:belongs-to,join:role_id=id"`

	base.AuditColumn
}
