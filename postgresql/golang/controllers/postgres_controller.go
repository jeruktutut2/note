package controllers

import (
	"net/http"
	modelrequests "note-golang-postgresql/models/requests"
	modelresponses "note-golang-postgresql/models/responses"
	"note-golang-postgresql/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostgresController interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type postgresController struct {
	PostgresService services.PostgresService
}

func NewPostgresController(postgresService services.PostgresService) PostgresController {
	return &postgresController{
		PostgresService: postgresService,
	}
}

func (controller *postgresController) Create(c echo.Context) error {
	var createRequest modelrequests.CreateRequest
	err := c.Bind(&createRequest)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusBadRequest, nil, []modelresponses.Error{{Field: "message", Message: "bad request"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}

	httpResponse := controller.PostgresService.Create(c.Request().Context(), createRequest)

	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *postgresController) Get(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}
	httpResponse := controller.PostgresService.Get(c.Request().Context(), id)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *postgresController) Update(c echo.Context) error {
	var updateRequest modelrequests.UpdateRequest
	err := c.Bind(&updateRequest)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}
	httpResponse := controller.PostgresService.Update(c.Request().Context(), updateRequest)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *postgresController) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}
	httpResponse := controller.PostgresService.Delete(c.Request().Context(), id)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}
