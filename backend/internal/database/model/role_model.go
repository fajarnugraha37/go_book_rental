package model

import (
	"backend/internal/database/model/base"

	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles,alias:roles"`

	ID          string       `bun:",pk"`
	IsSuperRole bool         `bun:"is_super_role,default:false"`
	Description string       `bun:"description"`
	Permissions []Permission `bun:"rel:has-many,join:id=role_id"`

	base.AuditColumn
}
