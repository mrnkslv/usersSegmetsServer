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

func (r *UserPostgres) AddUserToSlug(data models.AddSlugstoUser) ([]models.Slug, error) {
	//добавить транзакцию
	deleteQuery := fmt.Sprintf("delete from %s where user_id=$1 and slug_id=(select id from slugs where name=$2)", usersSlugsTable)
	for i := 0; i < len(data.OutdatedSlugs); i++ {
		r.db.QueryRow(deleteQuery, data.UserId, data.OutdatedSlugs[i].Name)
	}
	newSlugs := make([]models.Slug, len(data.NewSlugs))
	query := fmt.Sprintf("insert into %s (user_id,slug_id) values ((select id from users where id=$1), (select id from slugs where name=$2)) returning $2", usersSlugsTable)
	for i := 0; i < len(data.NewSlugs); i++ {
		row := r.db.QueryRow(query, data.UserId, data.NewSlugs[i].Name)
		row.Scan(&newSlugs[i].Name)
	}
	return newSlugs, nil
}

func (r *UserPostgres) GetActiveSlugsByID(userId int64) ([]models.Slug, error) {
	var ActiveSlugs []models.Slug
	query := fmt.Sprintf("select slug_id as id, name from %s inner join %s on slugs.id=users_slugs.slug_id where user_id=$1", slugsTable, usersSlugsTable)

	if err := r.db.Select(&ActiveSlugs, query, userId); err != nil {
		return nil, err
	}
	return ActiveSlugs, nil
}

func (r *UserPostgres) GetUserById(userId int64) (int64, error) {
	var getQueryTemplate = fmt.Sprintf("select id from %s where id =$1", usersTable)
	var user int64
	err := r.db.Get(&user, getQueryTemplate, userId)
	if err != nil {
		return 0, err
	}

	return user, nil
}
