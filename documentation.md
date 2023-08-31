# users-segments-service REST API Documentation

* [Introduction](#introduction)
* [Sql scripts](#sql-scripts)
* [Segments](#segments)
* [Users](#users)

## Introduction

This documentation provides a description of all available API handlers.

## Sql-scripts

To create the table use the next script:
```
create table users(
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
```
to fill in the users in users table use: 
```
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
values('Maria Romanova');
```
to delete tables use:
```
drop_table users_segments;
drop table users;
drop table segments;
```
## Segments

Used to create or delete segment

### POST api/segments

Used to create segment. Returns segment id.

Body
```
{
{
  "name":"AVITO_DISCOUNT_50"
}
}
```

Success response 
```
{
  "id": 1
}
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | segment name not valid 
| Internal Server Error | 500 | Server error. Possibly segment already exists

### DELETE api/segments

Used to delete segment. Returns segment id.

Body
```
{
{
  "name":"AVITO_DISCOUNT_50"
}
}
```

Success response 
```
{
  "id": 1
}
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | invalid characters in input
| Internal Server Error | 500 | Server error. Possibly segment's name incorrect.


## Users

Used to assign segments to user or delete them. Also used to get user's segments

### POST api/users

Create or delete user's segments

Body
```
{
  "new_segments":[{"name":"AVITO_DISCOUNT_30"},{"name":"AVITO_DISCOUNT_50"}],
  "outdated_segments":[{"name":"AVITO_DISCOUNT_80"}],
  "user_id":4
}
```

Success response ()
```
{
    "deleted_segments": [],
  "new_segments": [
    {
      "name": "AVITO_DISCOUNT_30"
    },
    {
        "name":"AVITO_DISCOUNT_50"
    }
  ]
}
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | invalid characters in input
| Internal Server Error | 500 | Server error

### GET api/users

Get all the user's segments

Form of request
```
localhost:8000/api/users/?id=4
```

Success response
```
[
  {
    "name": "AVITO_DISCOUNT_30"
  }
]
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Internal Server Error | 500 | Server error









