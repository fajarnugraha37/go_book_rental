package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"
)

type PermissionRepo interface {
	basev2.RepoDML[model.Permission]
	basev2.RepoDQL[model.Permission]
}
