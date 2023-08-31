# users-segments-service

* [Introduction](#introduction)
* [Quickstart](#quickstart)
* [Documentation](#documentation)
* [Dependencies](#dependencies)

## Introduction 

Users-segments-service is useful service for work with segments assigned to users. It has the following functionality: 
- Add new segment
- Delete segment
- Assign segments to user or delete them
- Get user's segments


## Quick_start
For quick start enter in the terminal command:
```
docker-compose up --build
```
In case of a database initialization error when launching the app container after the build, restart the container via docker commands or stop the process using (ctrl+c) and enter the command:
```
docker-compose up
```
To stop containers and remove containers,networks,volumes and images created by up: 
```
docker-compose down
```

## Documentation

Access full documentation on REST API [here](documentation.md).

## Dependencies

### build
docker, docker-compose

### database

PostgreSQL
