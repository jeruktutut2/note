package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	modelentities "timeout/models/entities"
)

type Test2Repository interface {
	Sleep(tx *sqlx.Tx, ctx context.Context, second int) (result string, err error)
	SleepWithDb(db *sqlx.DB, ctx context.Context, second int) (result string, err error)
	Create(tx *sqlx.Tx, ctx context.Context, test2 modelentities.Test2) (rowsAffected int64, err error)
	CreateWithDb(db *sqlx.DB, ctx context.Context, test2 modelentities.Test2) (rowsAffected int64, err error)
}

type Test2RepositoryImplementation struct {
}

func NewTest2Repository() Test2Repository {
	return &Test2RepositoryImplementation{}
}

func (repository *Test2RepositoryImplementation) Sleep(tx *sqlx.Tx, ctx context.Context, second int) (result string, err error) {
	err = tx.GetContext(ctx, &result, `SELECT PG_SLEEP($1);`, second)
	return
}

func (repository *Test2RepositoryImplementation) SleepWithDb(db *sqlx.DB, ctx context.Context, second int) (result string, err error) {
	err = db.GetContext(ctx, &result, `SELECT PG_SLEEP($1);`, second)
	return
}

func (repository Test2RepositoryImplementation) Create(tx *sqlx.Tx, ctx context.Context, test2 modelentities.Test2) (rowsAffected int64, err error) {
	result, err := tx.ExecContext(ctx, `INSERT INTO test2s(test) VALUES($1);`, test2.Test)
	if err != nil {
		return
	}
	return result.RowsAffected()
}

func (repository *Test2RepositoryImplementation) CreateWithDb(db *sqlx.DB, ctx context.Context, test2 modelentities.Test2) (rowsAffected int64, err error) {
	result, err := db.ExecContext(ctx, `INSERT INTO test2s(test) VALUES($1);`, test2.Test)
	if err != nil {
		return
	}
	return result.RowsAffected()
}
