package model

import (
	"backend/internal/database/model/base"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Rent struct {
	bun.BaseModel `bun:"table:rents,alias:rents"`

	ID         uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	UserID     uuid.UUID
	User       User
	BookID     uuid.UUID
	Book       Book
	RentAt     time.Time
	DueDate    time.Time
	ReturnAt   *time.Time
	IsReturned bool
	Amount     float64
	PaidAt     *time.Time
	IsPaid     bool

	base.AuditColumn
}
