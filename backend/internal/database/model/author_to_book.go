package model

import (
	"backend/internal/database/model/base"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AuthorToBook struct {
	bun.BaseModel `bun:"table:author_to_book,alias:author_to_book"`

	AuthorID uuid.UUID `bun:",pk"`
	Author   Author    `bun:"rel:belongs-to,join:author_id=id"`
	BookID   uuid.UUID `bun:",pk"`
	Book     Book      `bun:"rel:belongs-to,join:book_id=id"`

	base.AuditColumn
}
