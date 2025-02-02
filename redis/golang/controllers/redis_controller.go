package controllers

import (
	"net/http"
	modelrequests "note-golang-redis/models/requests"
	modelresponses "note-golang-redis/models/responses"
	"note-golang-redis/services"

	"github.com/labstack/echo/v4"
)

type RedisController interface {
	Set(c echo.Context) error
	Get(c echo.Context) error
	Del(c echo.Context) error
}

type RedisControllerImplementation struct {
	RedisService services.RedisService
}

func NewRedisController(redisService services.RedisService) RedisController {
	return &RedisControllerImplementation{
		RedisService: redisService,
	}
}

func (controller *RedisControllerImplementation) Set(c echo.Context) error {
	var createRequest modelrequests.CreateRequest
	err := c.Bind(&createRequest)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusBadRequest, nil, []modelresponses.Error{{Field: "message", Message: "bad request"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}
	httpResponse := controller.RedisService.Set(c.Request().Context(), createRequest)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *RedisControllerImplementation) Get(c echo.Context) error {
	httpResponse := controller.RedisService.Get(c.Request().Context(), "1")
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *RedisControllerImplementation) Del(c echo.Context) error {
	httpResponse := controller.RedisService.Del(c.Request().Context(), "1")
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}
