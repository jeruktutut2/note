package controllers

import (
	"net/http"
	modelrequests "note-golang-mysql/models/requests"
	modelresponses "note-golang-mysql/models/responses"
	"note-golang-mysql/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MysqlController interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type mysqlController struct {
	MysqlService services.MysqlService
}

func NewMysqlController(mysqlService services.MysqlService) MysqlController {
	return &mysqlController{
		MysqlService: mysqlService,
	}
}

func (controller *mysqlController) Create(c echo.Context) error {
	var createRequest modelrequests.CreateRequest
	err := c.Bind(&createRequest)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusBadRequest, nil, []modelresponses.Error{{Field: "message", Message: "bad request"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}

	httpResponse := controller.MysqlService.Create(c.Request().Context(), createRequest)

	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mysqlController) Get(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}
	httpResponse := controller.MysqlService.Get(c.Request().Context(), id)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mysqlController) Update(c echo.Context) error {
	var updateRequest modelrequests.UpdateRequest
	err := c.Bind(&updateRequest)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}
	httpResponse := controller.MysqlService.Update(c.Request().Context(), updateRequest)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}

func (controller *mysqlController) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		httpResponse := modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
	}
	httpResponse := controller.MysqlService.Delete(c.Request().Context(), id)
	return c.JSON(httpResponse.HttpStatusCode, httpResponse.Response)
}
