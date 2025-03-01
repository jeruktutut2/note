package services

import (
	"context"
	"database/sql"
	"net/http"
	modelentities "note-golang-postgresql/models/entities"
	modelrequests "note-golang-postgresql/models/requests"
	modelresponses "note-golang-postgresql/models/responses"
	"note-golang-postgresql/repositories"
	"note-golang-postgresql/utils"
)

type PostgresService interface {
	Create(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse)
	Get(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse)
	Update(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse)
	Delete(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse)
}

type postgresService struct {
	PostgresUtil       utils.PostgresUtil
	PostgresRepository repositories.PostgresRepository
}

func NewPostgresService(postgresUtil utils.PostgresUtil, postgresRepository repositories.PostgresRepository) PostgresService {
	return &postgresService{
		PostgresUtil:       postgresUtil,
		PostgresRepository: postgresRepository,
	}
}

func (service *postgresService) Create(ctx context.Context, createRequest modelrequests.CreateRequest) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.PostgresUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}

	defer func() {
		errCommitOrRollback := service.PostgresUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorHttpResponse()
		}
	}()

	var user modelentities.User
	user.Email = createRequest.Email
	user.Password = createRequest.Password
	id, err := service.PostgresRepository.Create(tx, ctx, user)
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}
	user.Id = id
	return modelresponses.SetDataHttpResponse(http.StatusCreated, user)
}

func (service *postgresService) Get(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse) {
	user, err := service.PostgresRepository.Get(service.PostgresUtil.GetDb(), ctx, id)
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}
	return modelresponses.SetDataHttpResponse(http.StatusCreated, user)
}

func (service *postgresService) Update(ctx context.Context, updateRequest modelrequests.UpdateRequest) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.PostgresUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}

	defer func() {
		errCommitOrRollback := service.PostgresUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorHttpResponse()
		}
	}()

	var user modelentities.User
	user.Id = updateRequest.Id
	user.Email = updateRequest.Email
	user.Password = updateRequest.Password
	rowsAffected, err := service.PostgresRepository.Update(tx, ctx, user)
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	} else if rowsAffected != 1 {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, modelresponses.MessageResponse{Message: "rows affected not one"})
	}
	return modelresponses.SetDataHttpResponse(http.StatusCreated, user)
}

func (service *postgresService) Delete(ctx context.Context, id int) (httpResponse modelresponses.HttpResponse) {
	tx, err := service.PostgresUtil.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	}

	defer func() {
		errCommitOrRollback := service.PostgresUtil.CommitOrRollback(tx, err)
		if errCommitOrRollback != nil {
			httpResponse = modelresponses.SetInternalServerErrorHttpResponse()
		}
	}()

	rowsAffected, err := service.PostgresRepository.Delete(tx, ctx, id)
	if err != nil {
		return modelresponses.SetInternalServerErrorHttpResponse()
	} else if rowsAffected != 1 {
		return modelresponses.SetHttpResponse(http.StatusInternalServerError, nil, modelresponses.MessageResponse{Message: "rows affected create mysql is not 1"})
	}
	return modelresponses.SetDataHttpResponse(http.StatusCreated, modelresponses.SetMessageHttpResponse("successfully delete user"))
}
