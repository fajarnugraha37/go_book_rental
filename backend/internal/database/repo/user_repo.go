package repo

import "backend/internal/database/model"

type UserRepo interface {
	baseRepo[UserFilter, model.User]
}

type UserFilter struct {
	CommonFilter
	Username  *string   `comparator:"eq" field:"username"`
	Usernames *[]string `comparator:"in" field:"username"`
}
