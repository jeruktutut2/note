# NOTE GOLANG REDIS

## library
    go get github.com/labstack/echo/v4
    go get github.com/redis/go-redis/v9

## env
    export REDIS_HOST=localhost
    export REDIS_PORT=6380
    export REDIS_DATABASE=0

## docker
    docker run --name project-redis -p 6380:6379 -d redis:latest