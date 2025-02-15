package model

import (
	"backend/internal/database/model/base"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PenaltyToRent struct {
	bun.BaseModel `bun:"table:penalty_to_rent,alias:penalty_to_rent"`

	PenaltyID uuid.UUID
	Penalty   Penalty
	RentID    uuid.UUID
	Rent      Rent

	base.AuditColumn
}
