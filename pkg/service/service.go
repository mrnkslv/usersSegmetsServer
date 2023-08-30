package service

import (
	"github.com/mrnkslv/user-segmentation-service/models"
	"github.com/mrnkslv/user-segmentation-service/pkg/repository"
)

type Slugger interface {
	CreateSlug(models.Slug) (int64, error)
	DeleteSlug(slug models.Slug) (int64, error)
}

type Users interface {
	AddUserToSlug(data models.AddSlugstoUser) (int64, error)
}

type Service struct {
	Slugger
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Slugger: NewSlugService(repos.Slugger),
		Users:   NewUserService(repos.User),
	}
}
