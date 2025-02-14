package services

import (
	"context"
	"database/sql"
	"net/http"
	modelentities "note-golang-mysql/models/entities"
	modelrequests "note-golang-mysql/models/requests"
	modelresponses "note-golang-mysql/models/responses"
	"note-golang-mysql/repositories"
	"note-golang-mysql/utils"
)

type MysqlService interface {
	Create(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse)
	Get(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse)
	Update(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse)
	Delete(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse)
}

type mysqlService struct {
	MysqlUtil       utils.MysqlUtil
	MysqlRepository repositories.MysqlRepository
}

func NewMysqlService(mysqlUtil utils.MysqlUtil, mysqlRepository repositories.MysqlRepository) MysqlService {
	return &mysqlService{
		MysqlUtil:       mysqlUtil,
		MysqlRepository: mysqlRepository,
	}
}

func (service *mysqlService) Create(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorHttpResponse()
		}
	}()

	var user modelentities.User
	user.Email = createRequest.Email
	user.Password = createRequest.Password
	rowsAffected, lastInsertedId, err := service.MysqlRepository.Create(tx, ctx, user)
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	} else if rowsAffected != 1 {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}
	user.Id = int(lastInsertedId)
	return modelresponses.SetDataHttpResponse(http.StatusCreated, user)
}

func (service *mysqlService) Get(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse) {
	user, err := service.MysqlRepository.Get(service.MysqlUtil.GetDb(), ctx, id)
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}
	return modelresponses.SetDataHttpResponse(http.StatusOK, user)
}

func (service *mysqlService) Update(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorHttpResponse()
		}
	}()

	var user modelentities.User
	user.Id = updateRequest.Id
	user.Email = updateRequest.Email
	user.Password = updateRequest.Password
	rowsAffected, err := service.MysqlRepository.Update(tx, ctx, user)
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	} else if rowsAffected != 1 {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}
	return modelresponses.SetDataHttpResponse(http.StatusOK, user)
}

func (service *mysqlService) Delete(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorHttpResponse()
		}
	}()

	rowsAffected, err := service.MysqlRepository.Delete(tx, ctx, id)
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	} else if rowsAffected != 1 {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}
	return modelresponses.SetMessageHttpResponse(http.StatusNoContent, "successfully delete user")
}
