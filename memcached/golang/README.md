# MEMCACHED

## library
    go get github.com/labstack/echo/v4
    go get github.com/bradfitz/gomemcache/memcache

## docker
    docker run --name project-memcached -p 11211:11211 -d memcached
    docker run -d --name memcached-container -p 11211:11211 memcached memcached -m 64 -vv
    -m 64: limiting memory to 64 MB.
    -vv: using verbose logging.
