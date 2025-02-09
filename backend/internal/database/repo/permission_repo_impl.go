package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type PermissionRepoImpl struct {
	db *bun.DB
	*basev2.RepoDMLImpl[model.Permission]
	*basev2.RepoDQLImpl[model.Permission]
}

func NewPermissionRepo(db *bun.DB) PermissionRepo {
	var _ PermissionRepo = (*PermissionRepoImpl)(nil)
	return &PermissionRepoImpl{
		db:          db,
		RepoDMLImpl: basev2.NewRepoDML[model.Permission](db),
		RepoDQLImpl: basev2.NewRepoDQL[model.Permission](db),
	}
}
