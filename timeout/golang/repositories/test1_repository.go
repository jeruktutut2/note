package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	modelentities "timeout/models/entities"
)

type Test1Repository interface {
	Sleep(tx *sqlx.Tx, ctx context.Context, second int) (result string, err error)
	SleepWithDb(db *sqlx.DB, ctx context.Context, second int) (result string, err error)
	Create(tx *sqlx.Tx, ctx context.Context, test1 modelentities.Test1) (rowsAffected int64, err error)
	CreateWithDb(db *sqlx.DB, ctx context.Context, test1 modelentities.Test1) (rowsAffected int64, err error)
}

type Test1RepositoryImplementation struct {
}

func NewTest1Repository() Test1Repository {
	return &Test1RepositoryImplementation{}
}

func (repository *Test1RepositoryImplementation) Sleep(tx *sqlx.Tx, ctx context.Context, second int) (result string, err error) {
	err = tx.GetContext(ctx, &result, `SELECT PG_SLEEP($1);`, second)
	return
}

func (repository *Test1RepositoryImplementation) SleepWithDb(db *sqlx.DB, ctx context.Context, second int) (result string, err error) {
	err = db.GetContext(ctx, &result, `SELECT PG_SLEEP($1);`, second)
	return
}

func (repository Test1RepositoryImplementation) Create(tx *sqlx.Tx, ctx context.Context, test1 modelentities.Test1) (rowsAffected int64, err error) {
	result, err := tx.ExecContext(ctx, `INSERT INTO test1s(test) VALUES($1);`, test1.Test)
	if err != nil {
		return
	}
	return result.RowsAffected()
}

func (repository *Test1RepositoryImplementation) CreateWithDb(db *sqlx.DB, ctx context.Context, test1 modelentities.Test1) (rowsAffected int64, err error) {
	result, err := db.ExecContext(ctx, `INSERT INTO test1s(test) VALUES($1);`, test1.Test)
	if err != nil {
		return
	}
	return result.RowsAffected()
}
