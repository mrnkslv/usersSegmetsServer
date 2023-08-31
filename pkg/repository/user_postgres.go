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

/*func (r *UserPostgres) AddUserToSegments(data models.AddSegmentstoUser) ([]models.Segment, []models.Segment, error) {


	newSegments := make([]models.Segment, len(data.NewSegments))
	query := fmt.Sprintf("insert into %s (user_id,segment_id) values ((select id from users where id=$1), (select id from segments where name=$2)) returning $2", usersSegmentsTable)
	for i := 0; i < len(data.NewSegments); i++ {
		row := r.db.QueryRow(query, data.UserId, data.NewSegments[i].Name)
		row.Scan(&newSegments[i].Name)
	}
	deletedSegments := make([]models.Segment, len(data.OutdatedSegments))
	deleteQuery := fmt.Sprintf("delete from %s where user_id=$1 and segment_id=(select id from segments where name=$2)", usersSegmentsTable)
	for i := 0; i < len(data.OutdatedSegments); i++ {
		row := r.db.QueryRow(deleteQuery, data.UserId, data.OutdatedSegments[i].Name)
		row.Scan(&deletedSegments[i].Name)
	}
	return newSegments, deletedSegments, nil
}*/

func (r *UserPostgres) AddUserToSegments(data models.AddSegmentstoUser) ([]models.Segment, []models.Segment, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	deletedSegments := make([]models.Segment, len(data.OutdatedSegments))
	deleteQuery := fmt.Sprintf("delete from %s where user_id=$1 and segment_id=(select id from segments where name=$2)returning $2", usersSegmentsTable)
	for i := 0; i < len(data.OutdatedSegments); i++ {
		row := tx.QueryRow(deleteQuery, data.UserId, data.OutdatedSegments[i].Name)
		err = row.Scan(&deletedSegments[i].Name)
		if err != nil {
			tx.Rollback()
			return nil, nil, err
		}
	}

	newSegments := make([]models.Segment, len(data.NewSegments))
	query := fmt.Sprintf("insert into %s (user_id,segment_id) values ((select id from users where id=$1), (select id from segments where name=$2)) returning $2", usersSegmentsTable)
	for i := 0; i < len(data.NewSegments); i++ {
		row := tx.QueryRow(query, data.UserId, data.NewSegments[i].Name)
		err = row.Scan(&newSegments[i].Name)
		if err != nil {
			tx.Rollback()
			return nil, nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	return newSegments, deletedSegments, nil
}

func (r *UserPostgres) GetActiveSegmentsByID(userId int64) ([]models.Segment, error) {
	var ActiveSegments []models.Segment
	query := fmt.Sprintf("select segment_id as id, name from %s inner join %s on segments.id=users_segments.segment_id where user_id=$1", segmentsTable, usersSegmentsTable)

	if err := r.db.Select(&ActiveSegments, query, userId); err != nil {
		return nil, err
	}
	return ActiveSegments, nil
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
