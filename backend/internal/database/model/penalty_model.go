package model

import (
	"backend/internal/database/model/base"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Penalty struct {
	bun.BaseModel `bun:"table:penalties,alias:penalties"`

	ID       uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	UserID   uuid.UUID
	User     User
	StartAt  time.Time
	EndAt    *time.Time
	Reason   string
	IsActive bool
	ActiveAt *time.Time
	Amount   float64
	PaidAt   *time.Time
	IsPaid   bool

	base.AuditColumn
}
