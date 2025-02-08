package repo

import (
	"backend/internal/database/model"

	"github.com/uptrace/bun"
)

type UserRepoImpl struct {
	*baseRepoImpl[UserFilter, model.User]
	db *bun.DB
}

func NewUserRepo(db *bun.DB) UserRepo {
	var _ UserRepo = (*UserRepoImpl)(nil)
	return &UserRepoImpl{
		db:           db,
		baseRepoImpl: newBaseRepo[UserFilter, model.User](db),
	}
}
