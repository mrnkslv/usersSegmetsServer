package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrnkslv/user-segmentation-service/models"
)

type Slugger interface {
	CreateSlug(slug models.Slug) (int64, error)
}

type User interface {
}

type Repository struct {
	Slugger
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Slugger: NewSlugPostgres(db),
	}
}
