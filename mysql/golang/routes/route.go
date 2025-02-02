package routes

import (
	"note-golang-mysql/controllers"

	"github.com/labstack/echo/v4"
)

func SetRoute(e *echo.Echo, controller controllers.MysqlController) {
	e.POST("/create", controller.Create)
	e.GET("/get", controller.Get)
	e.PUT("/update", controller.Update)
	e.DELETE("/delete", controller.Delete)
}
