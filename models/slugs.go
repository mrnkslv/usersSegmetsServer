package models

type Slug struct {
	Id   int64  `db:"id" json:"-"`
	Name string `db:"name" json:"name"`
}
