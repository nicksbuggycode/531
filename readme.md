This project closely follows the Udemy Backend Masterclass course (https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/)

We are using: 

-- Docker container to run postgres15-alpine
---- docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

-- dbdiagram.io to generate the database structure

-- go-migrate for db migrations
---- migrate create -ext sql -dir db/migration init_schema
---- Put dbdiagram.io sql commands into the "up" init schema file
---- In "down" schema file, drop all tables if exists