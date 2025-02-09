package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type UserToRoleRepoImpl struct {
	db *bun.DB
	*basev2.RepoDMLImpl[model.UserToRole]
	*basev2.RepoDQLImpl[model.UserToRole]
}

func NewUserToRoleRepo(db *bun.DB) UserToRoleRepo {
	var _ UserToRoleRepo = (*UserToRoleRepoImpl)(nil)
	return &UserToRoleRepoImpl{
		db:          db,
		RepoDMLImpl: basev2.NewRepoDML[model.UserToRole](db),
		RepoDQLImpl: basev2.NewRepoDQL[model.UserToRole](db),
	}
}
