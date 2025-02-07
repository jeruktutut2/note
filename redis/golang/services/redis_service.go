package services

import (
	"context"
	"encoding/json"
	"net/http"
	modelrequests "note-golang-redis/models/requests"
	modelresponses "note-golang-redis/models/responses"
	"note-golang-redis/utils"
	"time"
)

type RedisService interface {
	Set(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse)
	Get(ctx context.Context, key string) (httpResponse modelresponses.HttpResponse)
	Del(ctx context.Context, key string) (httpResponse modelresponses.HttpResponse)
}

type redisService struct {
	RedisUtil utils.RedisUtil
}

func NewRedisService(redisUtil utils.RedisUtil) RedisService {
	return &redisService{
		RedisUtil: redisUtil,
	}
}

func (service *redisService) Set(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse) {
	resultBytes, err := json.Marshal(createRequest)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}
	_, err = service.RedisUtil.Set(ctx, "1", string(resultBytes), time.Duration(10)*time.Second)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}
	return modelresponses.SetHttpResponse(http.StatusCreated, createRequest, []modelresponses.Error{})
}

func (service *redisService) Get(ctx context.Context, key string) (httpResponse modelresponses.HttpResponse) {
	result, err := service.RedisUtil.Get(ctx, key)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}
	var createRequest modelrequests.CreateRequest
	err = json.Unmarshal([]byte(result), &createRequest)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}
	return modelresponses.SetHttpResponse(http.StatusOK, createRequest, []modelresponses.Error{})
}

func (service *redisService) Del(ctx context.Context, key string) (httpResponse modelresponses.HttpResponse) {
	rowsAffected, err := service.RedisUtil.Del(ctx, key)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	} else if rowsAffected != 1 {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}
	return modelresponses.SetHttpResponse(http.StatusNoContent, nil, []modelresponses.Error{})
}
