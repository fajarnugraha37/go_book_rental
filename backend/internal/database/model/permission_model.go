package model

import (
	"backend/internal/database/model/base"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Permission struct {
	bun.BaseModel `bun:"table:roles,alias:roles"`
	ID            uuid.UUID         `bun:"id,pk"`
	RoleID        string            `bun:"role_id"`
	Resource      string            `bun:"resource,notnull"`
	Privilege     map[string]string `bun:"privilege,notnull,type:jsonb,default:'{}'"`
	Expression    *string           `bun:"Expression"`
	base.AuditColumn

	Role *Role `bun:"rel:belongs-to,join:role_id=id"`
}
