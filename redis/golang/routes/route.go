package routes

import (
	"note-golang-redis/controllers"

	"github.com/labstack/echo/v4"
)

func SetRedisRoute(e *echo.Echo, controller controllers.RedisController) {
	e.POST("/redis", controller.Set)
	e.GET("/redis", controller.Get)
	e.DELETE("/redis", controller.Del)
}
