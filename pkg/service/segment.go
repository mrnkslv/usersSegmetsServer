package service

import (
	"github.com/mrnkslv/user-segmentation-service/models"
	"github.com/mrnkslv/user-segmentation-service/pkg/repository"
)

type Segment struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *Segment {
	return &Segment{repo: repo}
}

func (s *Segment) CreateSegment(segment models.Segment) (int64, error) {
	return s.repo.CreateSegment(segment)
}
func (s *Segment) DeleteSegment(segment models.Segment) (int64, error) {
	return s.repo.DeleteSegment(segment)
}
