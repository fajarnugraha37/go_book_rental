package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"
)

type RoleRepo interface {
	basev2.RepoDML[model.Role]
	basev2.RepoDQL[model.Role]
}
