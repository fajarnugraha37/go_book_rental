package model

import (
	"backend/internal/database/model/base"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ bun.BeforeInsertHook = (*User)(nil)
var _ bun.BeforeUpdateHook = (*User)(nil)

type User struct {
	bun.BaseModel `bun:"table:users,alias:users"`

	ID uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	// Provider           string     `bun:"provider,default:'password'"`
	Username           string     `bun:"username,unique,notnull"`
	Password           *string    `bun:"password,type:text"`
	IsActive           bool       `bun:"is_active"`
	ActiveAt           *time.Time `bun:"active_at"`
	ActiveBy           *string    `bun:"active_by"`
	ConfirmationToken  *string    `bun:"confirmation_token,unique"`
	ConfirmationExpiry *time.Time `bun:"confirmation_expiry"`
	OTPSecret          *string    `bun:"otp_secret"`
	OTPExpiry          *time.Time `bun:"otp_expiry"`
	MagicLinkToken     *string    `bun:"magic_link_token,unique"`
	MagicLinkExpiry    *time.Time `bun:"magic_link_expiry"`

	base.AuditColumn

	UserAttributes []*UserAttribute `bun:"rel:has-many,join:id=user_id"`
}
