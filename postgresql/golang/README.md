# NOTE POSTGRESQL

## library
    go get github.com/labstack/echo/v4
    go get github.com/jmoiron/sqlx
    go get github.com/lib/pq
    go get github.com/stretchr/testify

## docker
    docker pull postgres
    docker pull postgres:13.16
    docker run --name project-postgres -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=project_users -p 5432:5432 -d postgres:13.16
    docker exdc -it project-postgres bash

## postgresql
    psql -h localhost -d project_users -U postgres -W
    \list \l
    \c project_users
    \dt

    CREATE DATABASE ecommercev2;
    \c ecommercev2
    \dt

    CREATE TABLE users (
  	    id SERIAL PRIMARY KEY,
  	    username varchar(50) NOT NULL UNIQUE,
  	    email varchar(100) NOT NULL UNIQUE,
  	    password varchar(100) NOT NULL,
  	    created_at bigint NOT NULL
    );

    # please don't use " in insert values, use ' instead, or error will accoured, There is a column named "username" in table "users", but it cannot be referenced from this part of the query.
    INSERT INTO users (id,username,email,password,created_at) VALUES (1,'username','email@email.com','$2a$10$MvEM5qcQFk39jC/3fYzJzOIy7M/xQiGv/PAkkoarCMgsx/rO0UaPG',1695095017);