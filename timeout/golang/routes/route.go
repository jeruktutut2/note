package routes

import (
	"github.com/labstack/echo/v4"
	"timeout/controllers"
	"timeout/middlewares"
)

func TestRoute(e *echo.Echo, testController controllers.TestController) {
	e.POST("/test-with-tx-3s", testController.Test1WithTx, middlewares.SetTimeout3Seconds)
	e.POST("/test-with-tx-60s", testController.Test1WithTx, middlewares.SetTimeout60Seconds)
	e.POST("/test-without-tx-3s", testController.Test1WithoutTx, middlewares.SetTimeout3Seconds)
	e.POST("/test-without-tx-60s", testController.Test1WithoutTx, middlewares.SetTimeout60Seconds)
}
