package repo

import "backend/internal/database/model"

type UserAttributeRepo interface {
	baseRepo[UserAttributeFilter, model.UserAttribute]
}

type UserAttributeFilter struct {
	CommonFilter
	UserID  *string   `comparator:"eq" field:"user_id"`
	UserIDs *[]string `comparator:"in" field:"user_id"`
	Key     *string   `comparator:"eq" field:"key"`
	Keys    *[]string `comparator:"in" field:"key"`
	Type    *string   `comparator:"eq" field:"type"`
	Types   *[]string `comparator:"in" field:"type"`
	Value   *string   `comparator:"eq" field:"value"`
	Values  *[]string `comparator:"in" field:"value"`

	Username  *string   `comparator:"eq" field:"user.username"`
	Usernames *[]string `comparator:"in" field:"user.username"`
}
