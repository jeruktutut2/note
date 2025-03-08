# REDIS

## curl test
    curl -i -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/redis/1
    curl -i -X POST \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/redis
    curl -i -X DELETE \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/redis