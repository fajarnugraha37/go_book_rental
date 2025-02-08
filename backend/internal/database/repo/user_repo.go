package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"
)

type UserRepo interface {
	basev2.RepoDML[model.User]
	basev2.RepoDQL[model.User]
}

// v1
// type UserRepo interface {
// 	baseRepo[UserFilter, model.User]
// }

// type UserFilter struct {
// 	CommonFilter
// 	Username  *string   `comparator:"eq" field:"username"`
// 	Usernames *[]string `comparator:"in" field:"username"`
// }
