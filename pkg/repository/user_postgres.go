package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrnkslv/user-segmentation-service/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) AddUserToSlug(data models.AddSlugstoUser) (int64, error) {
	var slug_id int64
	//добавить транзакцию
	deleteQuery := fmt.Sprintf("delete from %s where user_id=$1 and slug_id=(select id from slugs where name=$2);", usersSlugsTable)
	for i := 0; i < len(data.OutdatedSlugs); i++ {
		r.db.QueryRow(deleteQuery, data.UserId, data.OutdatedSlugs[i].Name)
	}
	query := fmt.Sprintf("insert into %s (user_id,slug_id) values ($1, (select id from slugs where name=$2)) returning slug_id;", usersSlugsTable)
	for i := 0; i < len(data.NewSlugs); i++ {
		row := r.db.QueryRow(query, data.UserId, data.NewSlugs[i].Name)
		if err := row.Scan(&slug_id); err != nil {
			return 0, err
		}
	}
	return slug_id, nil
}
