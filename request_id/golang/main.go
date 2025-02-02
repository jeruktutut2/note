package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"request_id/middlewares"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		requestId, ok := c.Request().Context().Value(middlewares.RequestIdKey).(string)
		if !ok {
			return c.String(http.StatusOK, "requestId")
		}
		return c.String(http.StatusOK, requestId)
	}, middlewares.SetRequestId)
	e.Logger.Fatal(e.Start(":80"))
}
