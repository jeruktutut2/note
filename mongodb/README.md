# MONGODB

## curl test
    curl -i -X POST \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/test1/insert-one
    curl -i -X POST \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/test1/insert-many
    curl -i -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/test1/find-one/1
    curl -i -X GET -H "Content-Type: application/json" http://localhost:8080/api/v1/test1/find
    curl -i -X PUT \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/test1/update-one
    curl -i -X PUT \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/test1/update-by-id/7
    curl -i -X DELETE \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/test1/delete-one/7
    curl -i -X DELETE \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/test1/delete-many