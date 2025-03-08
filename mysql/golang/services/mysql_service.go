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
	Create(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.Response)
	Get(ctx context.Context, id int) (httpResponse modelresponses.Response)
	Update(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.Response)
	Delete(ctx context.Context, id int) (httpResponse modelresponses.Response)
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

func (service *mysqlService) Create(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.Response) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorResponse()
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorResponse()
		}
	}()

	var user modelentities.User
	user.Email = createRequest.Email
	user.Password = createRequest.Password
	rowsAffected, lastInsertedId, err := service.MysqlRepository.Create(tx, ctx, user)
	if err != nil {
		return modelresponses.SetInternalServerErrorResponse()
	} else if rowsAffected != 1 {
		return modelresponses.SetInternalServerErrorResponse()
	}
	user.Id = int(lastInsertedId)
	return modelresponses.SetCreatedResponse(user)
}

func (service *mysqlService) Get(ctx context.Context, id int) (httpResponse modelresponses.Response) {
	user, err := service.MysqlRepository.Get(service.MysqlUtil.GetDb(), ctx, id)
	if err != nil {
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetOkResponse(user)
}

func (service *mysqlService) Update(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.Response) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorResponse()
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorResponse()
		}
	}()

	var user modelentities.User
	user.Id = updateRequest.Id
	user.Email = updateRequest.Email
	user.Password = updateRequest.Password
	rowsAffected, err := service.MysqlRepository.Update(tx, ctx, user)
	if err != nil {
		return modelresponses.SetInternalServerErrorResponse()
	} else if rowsAffected != 1 {
		return modelresponses.SetInternalServerErrorResponse()
	}
	return modelresponses.SetDataResponse(http.StatusOK, user)
}

func (service *mysqlService) Delete(ctx context.Context, id int) (httpResponse modelresponses.Response) {
	tx, err := service.MysqlUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorResponse()
	}

	defer func() {
		errCommitOrRollback := service.MysqlUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorResponse()
		}
	}()

	rowsAffected, err := service.MysqlRepository.Delete(tx, ctx, id)
	if err != nil {
		return modelresponses.SetInternalServerErrorResponse()
	} else if rowsAffected != 1 {
		return modelresponses.SetInternalServerErrorResponse()
	}
	// return modelresponses.SetMessageResponse(http.StatusNoContent, "successfully delete user")
	return modelresponses.SetNoContentResponse("successfully delete user")
}
