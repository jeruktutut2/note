# KAFKA

## curl test
    curl -i -X POST \
        -H "Content-Type: application/json" \
        -d '{"test": "test post"}' \
        http://localhost:8080/api/v1/message/send-message