package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"timeout/services"
)

type TestController interface {
	Test1WithTx(c echo.Context) error
	Test1WithoutTx(c echo.Context) error
}

type TestControllerImplementation struct {
	TestService services.TestService
}

func NewTestController(testService services.TestService) TestController {
	return &TestControllerImplementation{
		TestService: testService,
	}
}

func (controller *TestControllerImplementation) Test1WithTx(c echo.Context) error {
	result := controller.TestService.TestWithTx(c.Request().Context())
	return c.JSON(http.StatusOK, map[string]interface{}{
		"key": result,
	})
}

func (controller *TestControllerImplementation) Test1WithoutTx(c echo.Context) error {
	result := controller.TestService.TestWithoutTx(c.Request().Context())
	return c.JSON(http.StatusOK, map[string]interface{}{
		"key": result,
	})
}
