package routes

import (
	"note-golang-mongodb/controllers"
	"note-golang-mongodb/helpers"
	"note-golang-mongodb/repositories"
	"note-golang-mongodb/services"
	"note-golang-mongodb/utils"

	"github.com/labstack/echo/v4"
)

func SetRoute(e *echo.Echo, mongoUtil utils.MongoUtil, uuidHelper helpers.UuidHelper) {
	mongodbRepository := repositories.NewMongodbRepository()
	mongodbService := services.NewMongodbService(mongoUtil, uuidHelper, mongodbRepository)
	mongodbController := controllers.NewMongodbController(mongodbService)
	e.POST("/api/v1/test1/insert-one", mongodbController.InsertOne)
	e.POST("/api/v1/test1/insert-many", mongodbController.InsertMany)
	e.GET("/api/v1/test1/find-one/:id", mongodbController.FindOne)
	e.GET("/api/v1/test1/find", mongodbController.Find)
	e.PUT("/api/v1/test1/update-one", mongodbController.UpdateOne)
	e.PUT("/api/v1/test1/update-by-id/:id", mongodbController.UpdateById)
	e.DELETE("/api/v1/test1/delete-one/:id", mongodbController.DeleteOne)
	e.DELETE("/api/v1/test1/delete-many", mongodbController.DeleteMany)
}
