package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type UserAttributeRepoImpl struct {
	db *bun.DB
	*basev2.RepoDMLImpl[model.UserAttribute]
	*basev2.RepoDQLImpl[model.UserAttribute]
}

func NewUserAttributeRepo(db *bun.DB) UserAttributeRepo {
	var _ UserAttributeRepo = (*UserAttributeRepoImpl)(nil)
	return &UserAttributeRepoImpl{
		db:          db,
		RepoDMLImpl: basev2.NewRepoDML[model.UserAttribute](db),
		RepoDQLImpl: basev2.NewRepoDQL[model.UserAttribute](db),
	}
}
