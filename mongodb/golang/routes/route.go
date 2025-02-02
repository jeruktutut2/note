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
	e.POST("/insert-one", mongodbController.InsertOne)
	e.POST("/insert-many", mongodbController.InsertMany)
	e.GET("/find-one", mongodbController.FindOne)
	e.GET("/find", mongodbController.Find)
	e.PUT("/update-one", mongodbController.UpdateOne)
	e.PUT("/update-by-id/:id", mongodbController.UpdateById)
	e.DELETE("/delete-one/:id", mongodbController.DeleteOne)
	e.DELETE("/delete-many", mongodbController.DeleteMany)
}
