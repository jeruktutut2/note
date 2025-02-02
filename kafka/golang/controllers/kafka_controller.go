package controllers

import (
	"net/http"
	"note-golang-kafka/services"

	"github.com/labstack/echo/v4"
)

type KafkaController interface {
	SendMessage(c echo.Context) error
}

type KafkaControllerImplementation struct {
	KafkaService services.KafkaService
}

func NewKafkaController(kafkaService services.KafkaService) KafkaController {
	return &KafkaControllerImplementation{
		KafkaService: kafkaService,
	}
}

func (controller *KafkaControllerImplementation) SendMessage(c echo.Context) error {
	textMessage := c.QueryParam("message")
	response := controller.KafkaService.SendMessage(c.Request().Context(), textMessage)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"response": response,
	})
}
