package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type userToRoleRepoImpl struct {
	db *bun.DB
	basev2.RepoDML[model.UserToRole]
	basev2.RepoDQL[model.UserToRole]
}

func NewUserToRoleRepo(db *bun.DB) UserToRoleRepo {
	var _ UserToRoleRepo = (*userToRoleRepoImpl)(nil)
	return &userToRoleRepoImpl{
		db:      db,
		RepoDML: basev2.NewRepoDML[model.UserToRole](db),
		RepoDQL: basev2.NewRepoDQL[model.UserToRole](db),
	}
}
