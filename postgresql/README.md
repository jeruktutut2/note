# POSTGRESQL

## install library
to parse json
    brew install jq
to do benchmarking
    brew install wrk
    wrk --version

## change file execute
    chmod +x test.sh
    ./test.sh

## benchmark using wrk
    wrk -t10 -c100 -d30s http://localhost:8080/test1/25
    wrk -t5 -c10 -d60s http://localhost:8080/api/v1/test1/25
    wrk -t10 -c10 -d60s http://localhost:8080/api/v1/test1/25
    wrk -t1 -c1 -d60s http://localhost:8080/api/v1/test1/25

## postgres process
    check status DATABASE:

    SELECT datname, usename, application_name, client_addr, state, count(*) as connections
    FROM pg_stat_activity
    GROUP BY datname, usename, application_name, client_addr, state
    ORDER BY connections DESC;
    
    SELECT datname, usename, application_name, client_addr, state
    FROM pg_stat_activity;


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