package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	modelentities "timeout/models/entities"
)

type Test3Repository interface {
	Sleep(tx *sqlx.Tx, ctx context.Context, second int) (result string, err error)
	SleepWithDb(db *sqlx.DB, ctx context.Context, second int) (result string, err error)
	Create(tx *sqlx.Tx, ctx context.Context, test3 modelentities.Test3) (rowsAffected int64, err error)
	CreateWithDb(db *sqlx.DB, ctx context.Context, test3 modelentities.Test3) (rowsAffected int64, err error)
}

type Test3RepositoryImplementation struct {
}

func NewTest3Repository() Test3Repository {
	return &Test3RepositoryImplementation{}
}

func (repository *Test3RepositoryImplementation) Sleep(tx *sqlx.Tx, ctx context.Context, second int) (result string, err error) {
	err = tx.GetContext(ctx, &result, `SELECT PG_SLEEP($1);`, second)
	return
}

func (repository *Test3RepositoryImplementation) SleepWithDb(db *sqlx.DB, ctx context.Context, second int) (result string, err error) {
	err = db.GetContext(ctx, &result, `SELECT PG_SLEEP($1);`, second)
	return
}

func (repository Test3RepositoryImplementation) Create(tx *sqlx.Tx, ctx context.Context, test3 modelentities.Test3) (rowsAffected int64, err error) {
	result, err := tx.ExecContext(ctx, `INSERT INTO test3s(test) VALUES($1);`, test3.Test)
	if err != nil {
		return
	}
	return result.RowsAffected()
}

func (repository *Test3RepositoryImplementation) CreateWithDb(db *sqlx.DB, ctx context.Context, test3 modelentities.Test3) (rowsAffected int64, err error) {
	result, err := db.ExecContext(ctx, `INSERT INTO test3s(test) VALUES($1);`, test3.Test)
	if err != nil {
		return
	}
	return result.RowsAffected()
}
