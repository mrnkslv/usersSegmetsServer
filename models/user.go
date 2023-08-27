package models

type User struct {
	Id   int    `json:"-" db:"id"	`
	Name string `json:"name"`
}

type UserSlugs struct {
	Id     int `json:"-" db:"id"`
	UserId int `json:"user_id"`
	SlugId int `json:"slug_id"`
}
