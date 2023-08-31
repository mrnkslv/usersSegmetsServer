package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrnkslv/user-segmentation-service/models"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}
func (r *SegmentPostgres) CreateSegment(segment models.Segment) (int64, error) {
	var id int64
	query := fmt.Sprintf("insert into %s (name) values ($1) returning id", segmentsTable)
	row := r.db.QueryRow(query, segment.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SegmentPostgres) DeleteSegment(segment models.Segment) (int64, error) {
	var id int64
	query := fmt.Sprintf("delete from %s where name=$1 returning id", segmentsTable)
	row := r.db.QueryRow(query, segment.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
