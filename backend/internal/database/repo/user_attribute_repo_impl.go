package repo

import (
	"backend/internal/database/model"
	"backend/internal/database/repo/basev2"

	"github.com/uptrace/bun"
)

type userAttributeRepoImpl struct {
	db *bun.DB
	basev2.RepoDML[model.UserAttribute]
	basev2.RepoDQL[model.UserAttribute]
}

func NewUserAttributeRepo(db *bun.DB) UserAttributeRepo {
	var _ UserAttributeRepo = (*userAttributeRepoImpl)(nil)
	return &userAttributeRepoImpl{
		db:      db,
		RepoDML: basev2.NewRepoDML[model.UserAttribute](db),
		RepoDQL: basev2.NewRepoDQL[model.UserAttribute](db),
	}
}
