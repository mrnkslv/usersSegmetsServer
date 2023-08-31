package service

import (
	"github.com/mrnkslv/user-segmentation-service/models"
	"github.com/mrnkslv/user-segmentation-service/pkg/repository"
)

type Segments interface {
	CreateSegment(models.Segment) (int64, error)
	DeleteSegment(models.Segment) (int64, error)
}

type Users interface {
	AddUserToSegments(data models.AddSegmentstoUser) ([]models.Segment, []models.Segment, error)
	GetActiveSegmentsByID(userId int64) ([]models.Segment, error)
	Exists(userId int64) (bool, error)
}

type Service struct {
	Segments
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Segments: NewSegmentService(repos.Segment),
		Users:    NewUserService(repos.User),
	}
}
