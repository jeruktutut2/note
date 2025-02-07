package controllers

import (
	"net/http"
	modelrequests "note-golang-mongodb/models/requests"
	"note-golang-mongodb/services"

	"github.com/labstack/echo/v4"
)

type MongodbController interface {
	InsertOne(c echo.Context) error
	InsertMany(c echo.Context) error
	FindOne(c echo.Context) error
	Find(c echo.Context) error
	UpdateOne(c echo.Context) error
	UpdateById(c echo.Context) error
	DeleteOne(c echo.Context) error
	DeleteMany(c echo.Context) error
}

type mongodbController struct {
	MongodbService services.MongodbService
}

func NewMongodbController(mongodbService services.MongodbService) MongodbController {
	return &mongodbController{
		MongodbService: mongodbService,
	}
}

func (controller *mongodbController) InsertOne(c echo.Context) error {
	var createRequest modelrequests.CreateRequest
	err := c.Bind(&createRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"response": err.Error(),
		})
	}
	httpResponse := controller.MongodbService.InsertOne(c.Request().Context(), createRequest)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mongodbController) InsertMany(c echo.Context) error {
	var createRequests []modelrequests.CreateRequest
	err := c.Bind(&createRequests)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"response": err.Error(),
		})
	}
	httpResponse := controller.MongodbService.InsertMany(c.Request().Context(), createRequests)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mongodbController) FindOne(c echo.Context) error {
	email := c.QueryParam("email")
	httpResponse := controller.MongodbService.FindOne(c.Request().Context(), email)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mongodbController) Find(c echo.Context) error {
	email := c.QueryParam("email")
	httpResponse := controller.MongodbService.Find(c.Request().Context(), email)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mongodbController) UpdateOne(c echo.Context) error {
	var updateRequest modelrequests.UpdateRequest
	err := c.Bind(&updateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"response": err.Error(),
		})
	}
	httpResponse := controller.MongodbService.UpdateOne(c.Request().Context(), updateRequest)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mongodbController) UpdateById(c echo.Context) error {
	id := c.Param("id")
	httpResponse := controller.MongodbService.UpdateById(c.Request().Context(), id)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mongodbController) DeleteOne(c echo.Context) error {
	id := c.Param("id")
	httpResponse := controller.MongodbService.DeleteOne(c.Request().Context(), id)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mongodbController) DeleteMany(c echo.Context) error {
	httpResponse := controller.MongodbService.DeleteMany(c.Request().Context())
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}
