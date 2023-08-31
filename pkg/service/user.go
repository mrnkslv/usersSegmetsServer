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

func (s *User) AddUserToSegments(data models.AddSegmentstoUser) ([]models.Segment, []models.Segment, error) {
	return s.repo.AddUserToSegments(data)
}

func (s *User) GetActiveSegmentsByID(userId int64) ([]models.Segment, error) {
	return s.repo.GetActiveSegmentsByID(userId)
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
