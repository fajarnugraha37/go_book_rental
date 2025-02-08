package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type UserRepoImpl struct {
	db *bun.DB
	*basev2.RepoDMLImpl[model.User]
	*basev2.RepoDQLImpl[model.User]
}

func NewUserRepo(db *bun.DB) UserRepo {
	var _ UserRepo = (*UserRepoImpl)(nil)
	return &UserRepoImpl{
		db:          db,
		RepoDMLImpl: basev2.NewRepoDML[model.User](db),
		RepoDQLImpl: basev2.NewRepoDQL[model.User](db),
	}
}
