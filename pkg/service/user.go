package service

import (
	"github.com/mrnkslv/user-segmentation-service/models"
	"github.com/mrnkslv/user-segmentation-service/pkg/repository"
)

type User struct {
	repo repository.User
}

func NewUserService(repo repository.User) *User {
	return &User{repo: repo}
}

func (s *User) AddUserToSlug(data models.AddSlugstoUser) (int64, error) {
	return s.repo.AddUserToSlug(data)
}

func (s *User) GetActiveSlugsByID(userId int64) ([]models.Slug, error) {
	return s.repo.GetActiveSlugsByID(userId)
}
