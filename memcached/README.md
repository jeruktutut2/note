# MEMCACHED

## curl test
    curl -i -X POST \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/memcached
    curl -i -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/test1/memcached/1
    curl -i -X DELETE \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/memcached