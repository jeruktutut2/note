package services

import (
	"context"
	"fmt"
	"note-golang-mongodb/helpers"
	modelentities "note-golang-mongodb/models/entitites"
	modelrequests "note-golang-mongodb/models/requests"
	modelresponses "note-golang-mongodb/models/responses"
	"note-golang-mongodb/repositories"
	"note-golang-mongodb/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongodbService interface {
	// InsertOne(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse)
	// InsertMany(ctx context.Context, createRequests []modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse)
	// FindOne(ctx context.Context, email string) (httpResponse modelresponses.HttpResponse)
	// Find(ctx context.Context, email string) (httpResponse modelresponses.HttpResponse)
	// UpdateOne(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse)
	// UpdateById(ctx context.Context, id string) (httpResponse modelresponses.HttpResponse)
	// DeleteOne(ctx context.Context, id string) (httpResponse modelresponses.HttpResponse)
	// DeleteMany(ctx context.Context) (httpResponse modelresponses.HttpResponse)
	Create(ctx context.Context, createRequest modelrequests.CreateRequest) (response modelresponses.Response)
	Get(ctx context.Context, test string) (response modelresponses.Response)
	GetById(ctx context.Context, id string) (response modelresponses.Response)
	UpdateOne(ctx context.Context, updateRequest modelrequests.UpdateRequest) (response modelresponses.Response)
	UpdateById(ctx context.Context, updateRequest modelrequests.UpdateRequest) (response modelresponses.Response)
	DeleteOne(ctx context.Context, deleteRequest modelrequests.DeleteRequest) (response modelresponses.Response)
}

type mongodbService struct {
	MongoUtil         utils.MongoUtil
	UuidHelper        helpers.UuidHelper
	MongodbRepository repositories.MongodbRepository
}

func NewMongodbService(mongoUtil utils.MongoUtil, uuidHelper helpers.UuidHelper, mongodbRepository repositories.MongodbRepository) MongodbService {
	return &mongodbService{
		MongoUtil:         mongoUtil,
		UuidHelper:        uuidHelper,
		MongodbRepository: mongodbRepository,
	}
}

// func (service *mongodbService) InsertOne(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse) {
// 	id, err := service.UuidHelper.GenerateUuidV7()
// 	if err != nil {
// 		fmt.Println("error when generating uuidv7", err)
// 		return modelresponses.SetInternalServerErrorResponse()
// 	}
// 	var user modelentities.User
// 	user.Id = id
// 	user.Email = createRequest.Email
// 	user.Password = createRequest.Password
// 	err = service.MongodbRepository.InsertOne(service.MongoUtil.GetDb(), ctx, user)
// 	if err != nil {
// 		fmt.Println("error when inserting one:", err)
// 		return modelresponses.SetInternalServerErrorResponse()
// 	}
// 	return modelresponses.SetDataHttpResponse(http.StatusCreated, user)
// }

// func (service *mongodbService) InsertMany(ctx context.Context, createRequests []modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse) {
// 	var users []modelentities.User
// 	for _, createRequest := range createRequests {
// 		id, err := service.UuidHelper.GenerateUuidV7()
// 		if err != nil {
// 			fmt.Println("error when generating uuidv7:", err)
// 			return modelresponses.SetInternalServerErrorResponse()
// 		}
// 		var user modelentities.User
// 		user.Id = id
// 		user.Email = createRequest.Email
// 		user.Password = createRequest.Password
// 		users = append(users, user)
// 	}
// 	err := service.MongodbRepository.InsertMany(service.MongoUtil.GetDb(), ctx, users)
// 	if err != nil {
// 		fmt.Println("error when inserting many:", err)
// 		return modelresponses.SetInternalServerErrorResponse()
// 	}

// 	return modelresponses.SetDataHttpResponse(http.StatusCreated, users)
// }

// func (service *mongodbService) FindOne(ctx context.Context, email string) (httpResponse modelresponses.HttpResponse) {
// 	user, err := service.MongodbRepository.FindOne(service.MongoUtil.GetDb(), ctx, email)
// 	if err != nil {
// 		fmt.Println("error when finding one:", err)
// 		if err == mongo.ErrNoDocuments || err == mongo.ErrNilDocument {
// 			return modelresponses.SetBadRequestResponse("cannot user find by email")
// 		}
// 		return modelresponses.SetInternalServerErrorResponse()
// 	}
// 	return modelresponses.SetDataHttpResponse(http.StatusOK, user)
// }

// func (service *mongodbService) Find(ctx context.Context, email string) (httpResponse modelresponses.HttpResponse) {
// 	users, err := service.MongodbRepository.Find(service.MongoUtil.GetDb(), ctx, email)
// 	if err != nil {
// 		fmt.Println("error when finding by email:", err)
// 		return modelresponses.SetInternalServerErrorResponse()
// 	} else if len(users) < 1 {
// 		fmt.Println("users not found")
// 		return modelresponses.SetNotFoundHttpResponse("users not found")
// 	}
// 	return modelresponses.SetDataHttpResponse(http.StatusOK, users)
// }

// func (service *mongodbService) UpdateOne(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse) {
// 	var user modelentities.User
// 	user.Id = updateRequest.Id
// 	user.Email = updateRequest.Email
// 	user.Password = updateRequest.Password
// 	rowsAffected, err := service.MongodbRepository.UpdateOne(service.MongoUtil.GetDb(), ctx, user)
// 	if err != nil {
// 		fmt.Println("error when updating one:", err)
// 		return modelresponses.SetInternalServerErrorResponse()
// 	} else if rowsAffected != 1 {
// 		fmt.Println("rows affected not one")
// 		return modelresponses.SetInternalServerErrorResponse()
// 	}
// 	return modelresponses.SetDataHttpResponse(http.StatusOK, user)
// }

// func (service *mongodbService) UpdateById(ctx context.Context, id string) (httpResponse modelresponses.HttpResponse) {
// 	rowsAffected, err := service.MongodbRepository.UpdateById(service.MongoUtil.GetDb(), ctx, id)
// 	if err != nil {
// 		fmt.Println("error when updating by id:", err)
// 		return modelresponses.SetInternalServerErrorResponse()
// 	} else if rowsAffected != 1 {
// 		fmt.Println("rows affected not one")
// 		return modelresponses.SetInternalServerErrorResponse()
// 	}
// 	return modelresponses.SetMessageHttpResponse(http.StatusOK, "update by id")
// }

// func (service *mongodbService) DeleteOne(ctx context.Context, id string) (httpResponse modelresponses.HttpResponse) {
// 	deletedCount, err := service.MongodbRepository.DeleteOne(service.MongoUtil.GetDb(), ctx, id)
// 	if err != nil {
// 		fmt.Println("error when delete one:", err)
// 		return modelresponses.SetInternalServerErrorResponse()
// 	} else if deletedCount != 1 {
// 		fmt.Println("deleted count not one")
// 		return modelresponses.SetInternalServerErrorResponse()
// 	}
// 	return modelresponses.SetMessageHttpResponse(http.StatusNoContent, "successfully delete one")
// }

// func (service *mongodbService) DeleteMany(ctx context.Context) (httpResponse modelresponses.HttpResponse) {
// 	_, err := service.MongodbRepository.DeleteMany(service.MongoUtil.GetDb(), ctx)
// 	if err != nil {
// 		fmt.Println("error when delete one:", err)
// 		return modelresponses.SetInternalServerErrorResponse()
// 	}
// 	return modelresponses.SetMessageHttpResponse(http.StatusNoContent, "successfully delete many")
// }

func (service *mongodbService) Create(ctx context.Context, createRequest modelrequests.CreateRequest) (response modelresponses.Response) {
	var test1 modelentities.Test1
	test1.Id = primitive.NewObjectID()
	test1.Test = createRequest.Test
	err := service.MongodbRepository.Create(service.MongoUtil.GetDb(), ctx, test1)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetCreatedResponse(modelresponses.SetCreateResponse(test1))
}

func (service *mongodbService) Get(ctx context.Context, test string) (response modelresponses.Response) {
	test1s, err := service.MongodbRepository.Get(service.MongoUtil.GetDb(), ctx, test)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetOkResponse(modelresponses.SetGetResponses(test1s))
}

func (service *mongodbService) GetById(ctx context.Context, id string) (response modelresponses.Response) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	// fmt.Println("objectId:", objectId)
	test1, err := service.MongodbRepository.GetById(service.MongoUtil.GetDb(), ctx, objectId)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetOkResponse(modelresponses.SetGetResponse(test1))
}

func (service *mongodbService) UpdateOne(ctx context.Context, updateRequest modelrequests.UpdateRequest) (response modelresponses.Response) {
	id, err := primitive.ObjectIDFromHex(updateRequest.Id)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	var test1 modelentities.Test1
	test1.Id = id
	test1.Test = updateRequest.Test
	rowsAffected, err := service.MongodbRepository.UpdateOne(service.MongoUtil.GetDb(), ctx, test1)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	fmt.Println("rowsAffected:", rowsAffected)
	return modelresponses.SetOkResponse(modelresponses.SetUpdateResponse(test1))
}

func (service *mongodbService) UpdateById(ctx context.Context, updateRequest modelrequests.UpdateRequest) (response modelresponses.Response) {
	id, err := primitive.ObjectIDFromHex(updateRequest.Id)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}

	var test1 modelentities.Test1
	test1.Id = id
	test1.Test = updateRequest.Test
	rowsAffected, err := service.MongodbRepository.UpdateById(service.MongoUtil.GetDb(), ctx, test1)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	fmt.Println("rowsAffected:", rowsAffected)
	return modelresponses.SetOkResponse(modelresponses.SetUpdateResponse(test1))
}

func (service *mongodbService) DeleteOne(ctx context.Context, deleteRequest modelrequests.DeleteRequest) (response modelresponses.Response) {
	objectId, err := primitive.ObjectIDFromHex(deleteRequest.Id)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	rowsAffected, err := service.MongodbRepository.DeleteOne(service.MongoUtil.GetDb(), ctx, objectId)
	if err != nil {
		fmt.Println("err:", err)
		return modelresponses.SetInternalServerErrorResponse()
	}
	fmt.Println("rowsAffected:", rowsAffected)
	return modelresponses.SetNoContentResponse()
}
