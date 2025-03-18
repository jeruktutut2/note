package controllers

import (
	modelrequests "note-golang-mongodb/models/requests"
	modelresponses "note-golang-mongodb/models/responses"
	"note-golang-mongodb/services"

	"github.com/labstack/echo/v4"
)

type MongodbController interface {
	// InsertOne(c echo.Context) error
	// InsertMany(c echo.Context) error
	// FindOne(c echo.Context) error
	// Find(c echo.Context) error
	// UpdateOne(c echo.Context) error
	// UpdateById(c echo.Context) error
	// DeleteOne(c echo.Context) error
	// DeleteMany(c echo.Context) error
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetById(c echo.Context) error
	UpdateOne(c echo.Context) error
	UpdateById(c echo.Context) error
	DeleteOne(c echo.Context) error
}

type mongodbController struct {
	MongodbService services.MongodbService
}

func NewMongodbController(mongodbService services.MongodbService) MongodbController {
	return &mongodbController{
		MongodbService: mongodbService,
	}
}

// func (controller *mongodbController) InsertOne(c echo.Context) error {
// 	var createRequest modelrequests.CreateRequest
// 	err := c.Bind(&createRequest)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"response": err.Error(),
// 		})
// 	}
// 	httpResponse := controller.MongodbService.InsertOne(c.Request().Context(), createRequest)
// 	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
// }

// func (controller *mongodbController) InsertMany(c echo.Context) error {
// 	var createRequests []modelrequests.CreateRequest
// 	err := c.Bind(&createRequests)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"response": err.Error(),
// 		})
// 	}
// 	httpResponse := controller.MongodbService.InsertMany(c.Request().Context(), createRequests)
// 	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
// }

// func (controller *mongodbController) FindOne(c echo.Context) error {
// 	email := c.QueryParam("email")
// 	httpResponse := controller.MongodbService.FindOne(c.Request().Context(), email)
// 	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
// }

// func (controller *mongodbController) Find(c echo.Context) error {
// 	email := c.QueryParam("email")
// 	httpResponse := controller.MongodbService.Find(c.Request().Context(), email)
// 	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
// }

// func (controller *mongodbController) UpdateOne(c echo.Context) error {
// 	var updateRequest modelrequests.UpdateRequest
// 	err := c.Bind(&updateRequest)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"response": err.Error(),
// 		})
// 	}
// 	httpResponse := controller.MongodbService.UpdateOne(c.Request().Context(), updateRequest)
// 	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
// }

// func (controller *mongodbController) UpdateById(c echo.Context) error {
// 	id := c.Param("id")
// 	httpResponse := controller.MongodbService.UpdateById(c.Request().Context(), id)
// 	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
// }

// func (controller *mongodbController) DeleteOne(c echo.Context) error {
// 	id := c.Param("id")
// 	httpResponse := controller.MongodbService.DeleteOne(c.Request().Context(), id)
// 	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
// }

// func (controller *mongodbController) DeleteMany(c echo.Context) error {
// 	httpResponse := controller.MongodbService.DeleteMany(c.Request().Context())
// 	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
// }

func (controller *mongodbController) Create(c echo.Context) error {
	var createRequest modelrequests.CreateRequest
	err := c.Bind(&createRequest)
	if err != nil {
		response := modelresponses.SetBadRequestResponse(err.Error())
		return c.JSON(response.HttpStatusCode, response.BodyResponse)
	}
	response := controller.MongodbService.Create(c.Request().Context(), createRequest)
	return c.JSON(response.HttpStatusCode, response.BodyResponse)
}

func (controller *mongodbController) Get(c echo.Context) error {
	test := c.QueryParam("test")
	response := controller.MongodbService.Get(c.Request().Context(), test)
	return c.JSON(response.HttpStatusCode, response.BodyResponse)
}

func (controller *mongodbController) GetById(c echo.Context) error {
	id := c.Param("id")
	response := controller.MongodbService.GetById(c.Request().Context(), id)
	return c.JSON(response.HttpStatusCode, response.BodyResponse)
}

func (controller *mongodbController) UpdateOne(c echo.Context) error {
	var updateRequest modelrequests.UpdateRequest
	err := c.Bind(&updateRequest)
	if err != nil {
		response := modelresponses.SetBadRequestResponse(err.Error())
		return c.JSON(response.HttpStatusCode, response.BodyResponse)
	}
	response := controller.MongodbService.UpdateById(c.Request().Context(), updateRequest)
	return c.JSON(response.HttpStatusCode, response.BodyResponse)
}

func (controller *mongodbController) UpdateById(c echo.Context) error {
	var updateRequest modelrequests.UpdateRequest
	err := c.Bind(&updateRequest)
	if err != nil {
		response := modelresponses.SetBadRequestResponse(err.Error())
		return c.JSON(response.HttpStatusCode, response.BodyResponse)
	}
	response := controller.MongodbService.UpdateById(c.Request().Context(), updateRequest)
	return c.JSON(response.HttpStatusCode, response.BodyResponse)
}

func (controller *mongodbController) DeleteOne(c echo.Context) error {
	var deleteRequest modelrequests.DeleteRequest
	err := c.Bind(&deleteRequest)
	if err != nil {
		response := modelresponses.SetBadRequestResponse(err.Error())
		return c.JSON(response.HttpStatusCode, response.BodyResponse)
	}
	response := controller.MongodbService.DeleteOne(c.Request().Context(), deleteRequest)
	return c.JSON(response.HttpStatusCode, response.BodyResponse)
}
