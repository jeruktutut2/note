package routes

import (
	"note-golang-memcached/controllers"
	"note-golang-memcached/services"
	"note-golang-memcached/utils"

	"github.com/labstack/echo/v4"
)

func SetMemcachedRoute(e *echo.Echo, memcachedUtil utils.MemcachedUtil) {
	memcachedService := services.NewMemcachedService(memcachedUtil)
	memcachedController := controllers.NewMemcachedController(memcachedService)
	e.POST("/set", memcachedController.Set)
	e.GET("/get", memcachedController.Get)
	e.DELETE("/delete", memcachedController.Delete)
}
