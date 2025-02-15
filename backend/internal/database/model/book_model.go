package model

import (
	"backend/internal/database/model/base"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Book struct {
	bun.BaseModel `bun:"table:books,alias:books"`

	ID                uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	PublisherID       uuid.UUID
	Publisher         Publisher
	ISBN              string
	Title             string
	Description       string
	Year              time.Time
	Language          string
	Quantity          int
	AvailableQuantity int
	RentQuantity      int
	InProcessQuantity int
	Authors           []Author

	base.AuditColumn
}
