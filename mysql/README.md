# MYSQL

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

## check mysql process
    SHOW PROCESSLIST;

## curl test
    curl -i -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/test1/25
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