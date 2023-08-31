package service

import (
	"database/sql"

	"github.com/mrnkslv/user-segmentation-service/models"
	"github.com/mrnkslv/user-segmentation-service/pkg/repository"
	"github.com/pkg/errors"
)

type User struct {
	repo repository.User
}

func NewUserService(repo repository.User) *User {
	return &User{repo: repo}
}

func (s *User) AddUserToSlug(data models.AddSlugstoUser) ([]models.Slug, error) {
	return s.repo.AddUserToSlug(data)
}

func (s *User) GetActiveSlugsByID(userId int64) ([]models.Slug, error) {
	return s.repo.GetActiveSlugsByID(userId)
}

func (s *User) Exists(userId int64) (bool, error) {
	_, err := s.repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.Wrap(err, "user not exists")
		}
		return false, errors.Wrap(err, "can't check if user exists")
	}

	return true, nil
}
