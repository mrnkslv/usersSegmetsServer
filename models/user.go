package models

type User struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" binding:"required"`
}

type UserSlugs struct {
	Id     int64 `json:"-" db:"id"`
	UserId int64 `json:"user_id"`
	SlugId int64 `json:"slug_id"`
}

type AddSlugstoUser struct {
	NewSlugs      []Slug `json:"new_slugs"`
	OutdatedSlugs []Slug `json:"outdated_slugs"`
	UserId        int64  `json:"user_id"`
}

/*create table users(
    id bigserial primary key,
    name varchar(128) not null unique
);

create table slugs (
    id bigserial primary key,
    name varchar(128) not null unique
);

create table users_slugs (
	id bigserial,
	user_id bigint not null,
	slug_id bigint not null,
	primary key(user_id,slug_id)
);

insert into users (name)
values('Anna Ivanova');
insert into users (name)
values('Elena Stepanova');
insert into users (name)
values('Maxim Korolev');
insert into users (name)
values('Egor Lipin');
insert into users (name)
values('Sergey Sidorov');
insert into users (name)
values('Maria Romanova');*/
