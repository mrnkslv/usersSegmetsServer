package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrnkslv/user-segmentation-service/models"
)

type SlugPostgres struct {
	db *sqlx.DB
}

func NewSlugPostgres(db *sqlx.DB) *SlugPostgres {
	return &SlugPostgres{db: db}
}
func (r *SlugPostgres) CreateSlug(slug models.Slug) (int64, error) {
	var id int64
	query := fmt.Sprintf("insert into %s (name) values ($1) returning id", slugsTable)
	row := r.db.QueryRow(query, slug.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SlugPostgres) DeleteSlug(slug models.Slug) (int64, error) {
	var id int64
	query := fmt.Sprintf("delete from %s where name=$1 returning id", slugsTable)
	row := r.db.QueryRow(query, slug.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
