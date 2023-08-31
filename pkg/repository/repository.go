package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrnkslv/user-segmentation-service/models"
)

type Segment interface {
	CreateSegment(slug models.Segment) (int64, error)
	DeleteSegment(slug models.Segment) (int64, error)
}

type User interface {
	AddUserToSegments(data models.AddSegmentstoUser) ([]models.Segment, []models.Segment, error)
	GetActiveSegmentsByID(userId int64) ([]models.Segment, error)
	GetUserById(userId int64) (int64, error)
}

type Repository struct {
	Segment
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segment: NewSegmentPostgres(db),
		User:    NewUserPostgres(db),
	}
}
