package services

import (
	"context"
	"fmt"
	"net/http"
	"note-golang-mongodb/helpers"
	modelentities "note-golang-mongodb/models/entitites"
	modelrequests "note-golang-mongodb/models/requests"
	modelresponses "note-golang-mongodb/models/responses"
	"note-golang-mongodb/repositories"
	"note-golang-mongodb/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbService interface {
	InsertOne(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse)
	InsertMany(ctx context.Context, createRequests []modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse)
	FindOne(ctx context.Context, email string) (httpResponse modelresponses.HttpResponse)
	Find(ctx context.Context, email string) (httpResponse modelresponses.HttpResponse)
	UpdateOne(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse)
	UpdateById(ctx context.Context, id string) (httpResponse modelresponses.HttpResponse)
	DeleteOne(ctx context.Context, id string) (httpResponse modelresponses.HttpResponse)
	DeleteMany(ctx context.Context) (httpResponse modelresponses.HttpResponse)
}

type MongodbServiceImplementation struct {
	MongoUtil         utils.MongoUtil
	UuidHelper        helpers.UuidHelper
	MongodbRepository repositories.MongodbRepository
}

func NewMongodbService(mongoUtil utils.MongoUtil, uuidHelper helpers.UuidHelper, mongodbRepository repositories.MongodbRepository) MongodbService {
	return &MongodbServiceImplementation{
		MongoUtil:         mongoUtil,
		UuidHelper:        uuidHelper,
		MongodbRepository: mongodbRepository,
	}
}

func (service *MongodbServiceImplementation) InsertOne(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse) {
	id, err := service.UuidHelper.GenerateUuidV7()
	if err != nil {
		fmt.Println("error when generating uuidv7", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	var user modelentities.User
	user.Id = id
	user.Email = createRequest.Email
	user.Password = createRequest.Password
	err = service.MongodbRepository.InsertOne(service.MongoUtil.GetDb(), ctx, user)
	if err != nil {
		fmt.Println("error when inserting one:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetHttpResponse(http.StatusCreated, user, []modelresponses.Error{})
}

func (service *MongodbServiceImplementation) InsertMany(ctx context.Context, createRequests []modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse) {
	var users []modelentities.User
	for _, createRequest := range createRequests {
		id, err := service.UuidHelper.GenerateUuidV7()
		if err != nil {
			fmt.Println("error when generating uuidv7:", err)
			return modelresponses.SetInternalServerErrorResponse()
		}
		var user modelentities.User
		user.Id = id
		user.Email = createRequest.Email
		user.Password = createRequest.Password
		users = append(users, user)
	}
	err := service.MongodbRepository.InsertMany(service.MongoUtil.GetDb(), ctx, users)
	if err != nil {
		fmt.Println("error when inserting many:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}

	return modelresponses.SetHttpResponse(http.StatusCreated, users, []modelresponses.Error{})
}

func (service *MongodbServiceImplementation) FindOne(ctx context.Context, email string) (httpResponse modelresponses.HttpResponse) {
	user, err := service.MongodbRepository.FindOne(service.MongoUtil.GetDb(), ctx, email)
	if err != nil {
		fmt.Println("error when finding one:", err)
		if err == mongo.ErrNoDocuments || err == mongo.ErrNilDocument {
			return modelresponses.SetBadRequestResponse("message", "cannot user find by email")
		}
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetHttpResponse(http.StatusOK, user, []modelresponses.Error{})
}

func (service *MongodbServiceImplementation) Find(ctx context.Context, email string) (httpResponse modelresponses.HttpResponse) {
	users, err := service.MongodbRepository.Find(service.MongoUtil.GetDb(), ctx, email)
	if err != nil {
		fmt.Println("error when finding by email:", err)
		return modelresponses.SetInternalServerErrorResponse()
	} else if len(users) < 1 {
		fmt.Println("users not found")
		return modelresponses.SetNotFoundHttpResponse("message", "users not found")
	}
	return modelresponses.SetHttpResponse(http.StatusOK, users, []modelresponses.Error{})
}

func (service *MongodbServiceImplementation) UpdateOne(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse) {
	var user modelentities.User
	user.Id = updateRequest.Id
	user.Email = updateRequest.Email
	user.Password = updateRequest.Password
	rowsAffected, err := service.MongodbRepository.UpdateOne(service.MongoUtil.GetDb(), ctx, user)
	if err != nil {
		fmt.Println("error when updating one:", err)
		return modelresponses.SetInternalServerErrorResponse()
	} else if rowsAffected != 1 {
		fmt.Println("rows affected not one")
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetHttpResponse(http.StatusOK, user, []modelresponses.Error{})
}

func (service *MongodbServiceImplementation) UpdateById(ctx context.Context, id string) (httpResponse modelresponses.HttpResponse) {
	rowsAffected, err := service.MongodbRepository.UpdateById(service.MongoUtil.GetDb(), ctx, id)
	if err != nil {
		fmt.Println("error when updating by id:", err)
		return modelresponses.SetInternalServerErrorResponse()
	} else if rowsAffected != 1 {
		fmt.Println("rows affected not one")
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetMessageHttpResponse(http.StatusOK, "update by id")
}

func (service *MongodbServiceImplementation) DeleteOne(ctx context.Context, id string) (httpResponse modelresponses.HttpResponse) {
	deletedCount, err := service.MongodbRepository.DeleteOne(service.MongoUtil.GetDb(), ctx, id)
	if err != nil {
		fmt.Println("error when delete one:", err)
		return modelresponses.SetInternalServerErrorResponse()
	} else if deletedCount != 1 {
		fmt.Println("deleted count not one")
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetMessageHttpResponse(http.StatusNoContent, "successfully delete one")
}

func (service *MongodbServiceImplementation) DeleteMany(ctx context.Context) (httpResponse modelresponses.HttpResponse) {
	_, err := service.MongodbRepository.DeleteMany(service.MongoUtil.GetDb(), ctx)
	if err != nil {
		fmt.Println("error when delete one:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetMessageHttpResponse(http.StatusNoContent, "successfully delete many")
}
