package model

import (
	"backend/internal/database/model/base"
	"backend/pkg/helper"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ bun.BeforeAppendModelHook = (*User)(nil)

type User struct {
	bun.BaseModel `bun:"table:users,alias:users"`

	ID                 uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Provider           string     `bun:"provider,default:'password'"`
	Username           string     `bun:"username,unique,notnull"`
	Password           *string    `bun:"password,type:text"`
	IsActive           bool       `bun:"is_active"`
	ActiveAt           *time.Time `bun:"active_at,type:timestamptz"`
	ActiveBy           *string    `bun:"active_by"`
	ConfirmationToken  *string    `bun:"confirmation_token,unique"`
	ConfirmationExpiry *time.Time `bun:"confirmation_expiry,type:timestamptz"`
	OTPSecret          *string    `bun:"otp_secret"`
	OTPExpiry          *time.Time `bun:"otp_expiry,type:timestamptz"`
	MagicLinkToken     *string    `bun:"magic_link_token,unique"`
	MagicLinkExpiry    *time.Time `bun:"magic_link_expiry,type:timestamptz"`

	base.AuditColumn

	UserAttributes []*UserAttribute `bun:"rel:has-many,join:id=user_id"`
}

func (m *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = helper.ToPtr(time.Now())
		if m.CreatedBy == nil {
			m.CreatedBy = helper.ToPtr("system@intranet")
		}
	case *bun.UpdateQuery:
		m.UpdatedAt = helper.ToPtr(time.Now())
		if m.UpdatedBy == nil {
			m.UpdatedBy = helper.ToPtr("system@intranet")
		}
	}
	return nil
}
