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

type MysqlServiceImplementation struct {
	MysqlUtil       utils.MysqlUtil
	MysqlRepository repositories.MysqlRepository
}

func NewMysqlService(mysqlUtil utils.MysqlUtil, mysqlRepository repositories.MysqlRepository) MysqlService {
	return &MysqlServiceImplementation{
		MysqlUtil:       mysqlUtil,
		MysqlRepository: mysqlRepository,
	}
}

func (service *MysqlServiceImplementation) Create(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		}
	}()

	var user modelentities.User
	user.Email = createRequest.Email
	user.Password = createRequest.Password
	rowsAffected, lastInsertedId, err := service.MysqlRepository.Create(tx, ctx, user)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	} else if rowsAffected != 1 {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "rows affected create mysql is not 1"}})
	}
	user.Id = int(lastInsertedId)
	return modelresponses.SetHttpResponse(http.StatusCreated, user, []modelresponses.Error{})
}

func (service *MysqlServiceImplementation) Get(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse) {
	user, err := service.MysqlRepository.Get(service.MysqlUtil.GetDb(), ctx, id)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}
	return modelresponses.SetHttpResponse(http.StatusCreated, user, []modelresponses.Error{})
}

func (service *MysqlServiceImplementation) Update(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		}
	}()

	var user modelentities.User
	user.Id = updateRequest.Id
	user.Email = updateRequest.Email
	user.Password = updateRequest.Password
	rowsAffected, err := service.MysqlRepository.Update(tx, ctx, user)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	} else if rowsAffected != 1 {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "rows affected create mysql is not 1"}})
	}
	return modelresponses.SetHttpResponse(http.StatusCreated, user, []modelresponses.Error{})
}

func (service *MysqlServiceImplementation) Delete(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
		}
	}()

	rowsAffected, err := service.MysqlRepository.Delete(tx, ctx, id)
	if err != nil {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "internal server error"}})
	} else if rowsAffected != 1 {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, []modelresponses.Error{{Field: "message", Message: "rows affected create mysql is not 1"}})
	}
	return modelresponses.SetHttpResponse(http.StatusCreated, modelresponses.SetMessageHttpResponse("successfully delete user"), []modelresponses.Error{})
}
