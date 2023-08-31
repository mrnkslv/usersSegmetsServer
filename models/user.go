package models

type User struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" binding:"required"`
}

type UsersSegments struct {
	Id        int64 `json:"-" db:"id"`
	UserId    int64 `json:"user_id"`
	SegmentId int64 `json:"segment_id"`
}

type AddSegmentstoUser struct {
	NewSegments      []Segment `json:"new_segments"`
	OutdatedSegments []Segment `json:"outdated_segments"`
	UserId           int64     `json:"user_id"`
}
