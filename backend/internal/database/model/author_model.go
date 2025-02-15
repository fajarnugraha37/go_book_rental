package model

import (
	"backend/internal/database/model/base"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Author struct {
	bun.BaseModel `bun:"table:authors,alias:authors"`

	ID    uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Name  string
	Bio   *string
	Phone *string
	Email *string
	Books []Book

	base.AuditColumn
}
