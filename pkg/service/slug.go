package service

import (
	"github.com/mrnkslv/user-segmentation-service/models"
	"github.com/mrnkslv/user-segmentation-service/pkg/repository"
)

type Slug struct {
	repo repository.Slugger
}

func NewSlugService(repo repository.Slugger) *Slug {
	return &Slug{repo: repo}
}

func (s *Slug) CreateSlug(slug models.Slug) (int64, error) {
	return s.repo.CreateSlug(slug)
}
