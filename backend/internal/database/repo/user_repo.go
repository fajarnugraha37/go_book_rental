package repo

import "backend/internal/database/model"

type UserRepo interface {
	baseRepo[UserFilter, model.User]
}

type UserFilter struct {
	AuditFilter
	Username *string `comparator:"like" field:"username"`
}
