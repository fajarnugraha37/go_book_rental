package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type permissionRepoImpl struct {
	db *bun.DB
	basev2.RepoDML[model.Permission]
	basev2.RepoDQL[model.Permission]
}

func NewPermissionRepo(db *bun.DB) PermissionRepo {
	var _ PermissionRepo = (*permissionRepoImpl)(nil)
	return &permissionRepoImpl{
		db:      db,
		RepoDML: basev2.NewRepoDML[model.Permission](db),
		RepoDQL: basev2.NewRepoDQL[model.Permission](db),
	}
}
