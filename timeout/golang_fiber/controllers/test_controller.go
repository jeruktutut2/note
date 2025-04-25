package controllers

import (
	"context"
	"fmt"
	"net/http"
	"note-golang-fiber-timeout/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TestController interface {
	Test1WithTx(c *fiber.Ctx) error
	Test1WithoutTx(c *fiber.Ctx) error
}

type testController struct {
	TestService services.TestService
}

func NewTestController(testService services.TestService) TestController {
	return &testController{
		TestService: testService,
	}
}

func (controller *testController) Test1WithTx(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), time.Duration(7)*time.Second)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			c.Status(http.StatusOK).JSON(map[string]string{
				"response": "mantap",
			})
		}
	}()

	response := controller.TestService.TestWithTx(ctx)
	return c.Status(http.StatusOK).JSON(map[string]string{
		"response": response,
	})
}

func (controller *testController) Test1WithoutTx(c *fiber.Ctx) error {
	response := controller.TestService.TestWithoutTx(c.Context())
	return c.Status(http.StatusOK).JSON(map[string]string{
		"response": response,
	})
}
