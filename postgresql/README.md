# POSTGRESQL

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

## curl test
    curl -i -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/test1/1
    curl -i -X POST \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/test1
    curl -i -X PUT \
        -H "Content-Type: application/json" \
        -d '{"id": 6, "test": "test put 6"}' \
        http://localhost:8080/api/v1/test1
    curl -i -X DELETE \
        -H "Content-Type: application/json" \
        -d '{"id": 6}' \
        http://localhost:8080/api/v1/test1