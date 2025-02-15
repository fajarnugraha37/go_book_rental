package model

import (
	"backend/internal/database/model/base"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserToRole struct {
	bun.BaseModel `bun:"table:user_to_role,alias:user_to_role"`

	UserID uuid.UUID `bun:",pk"`
	User   User      `bun:"rel:belongs-to,join:user_id=id"`
	RoleID string    `bun:",pk"`
	Role   Role      `bun:"rel:belongs-to,join:role_id=id"`

	base.AuditColumn
}
