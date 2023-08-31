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

/*create table users(
    id bigserial primary key,
    name varchar(128) not null unique
);

create table segments (
    id bigserial primary key,
    name varchar(128) not null unique
);

create table users_segments (
	id bigserial,
	user_id bigint not null,
	segment_id bigint not null,
	primary key(user_id,segment_id)
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
