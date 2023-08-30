package models

type User struct {
	Id   int64  `json:"id" db:"id"	`
	Name string `json:"name" binding:"required"`
}

type UserSlugs struct {
	Id     int64 `json:"-" db:"id"`
	UserId int64 `json:"user_id"`
	SlugId int64 `json:"slug_id"`
}

/*create table users(
    id bigserial primary key,
    name varchar(128) not null
);

create table slugs (
    id bigserial primary key,
    name varchar(128) not null
);

create table users_slugs (
	id bigserial primary key,
	user_id bigint not null,
	slug_id bigint not null
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
