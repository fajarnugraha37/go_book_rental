package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type userRepoImpl struct {
	db *bun.DB
	basev2.RepoDML[model.User]
	basev2.RepoDQL[model.User]
}

func NewUserRepo(db *bun.DB) UserRepo {
	var _ UserRepo = (*userRepoImpl)(nil)
	return &userRepoImpl{
		db:      db,
		RepoDML: basev2.NewRepoDML[model.User](db),
		RepoDQL: basev2.NewRepoDQL[model.User](db),
	}
}
