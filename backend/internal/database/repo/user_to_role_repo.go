package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"
)

type UserToRoleRepo interface {
	basev2.RepoDML[model.UserToRole]
	basev2.RepoDQL[model.UserToRole]
}
