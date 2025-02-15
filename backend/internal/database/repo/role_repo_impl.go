package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type roleRepoImpl struct {
	db *bun.DB
	basev2.RepoDML[model.Role]
	basev2.RepoDQL[model.Role]
}

func NewRoleRepo(db *bun.DB) RoleRepo {
	var _ RoleRepo = (*roleRepoImpl)(nil)
	return &roleRepoImpl{
		db:      db,
		RepoDML: basev2.NewRepoDML[model.Role](db),
		RepoDQL: basev2.NewRepoDQL[model.Role](db),
	}
}
