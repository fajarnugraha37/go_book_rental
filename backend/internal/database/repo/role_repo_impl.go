package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type RoleRepoImpl struct {
	db *bun.DB
	*basev2.RepoDMLImpl[model.Role]
	*basev2.RepoDQLImpl[model.Role]
}

func NewRoleRepo(db *bun.DB) RoleRepo {
	var _ RoleRepo = (*RoleRepoImpl)(nil)
	return &RoleRepoImpl{
		db:          db,
		RepoDMLImpl: basev2.NewRepoDML[model.Role](db),
		RepoDQLImpl: basev2.NewRepoDQL[model.Role](db),
	}
}
