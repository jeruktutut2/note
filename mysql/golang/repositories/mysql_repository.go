package repositories

import (
	"context"
	modelentities "note-golang-mysql/models/entities"

	"github.com/jmoiron/sqlx"
)

type MysqlRepository interface {
	Create(tx *sqlx.Tx, ctx context.Context, user modelentities.User) (rowsAffected int64, lastInsertedId int64, err error)
	Get(db *sqlx.DB, ctx context.Context, id int) (user modelentities.User, err error)
	Update(tx *sqlx.Tx, ctx context.Context, user modelentities.User) (rowsAffected int64, err error)
	Delete(tx *sqlx.Tx, ctx context.Context, id int) (rowsAffected int64, err error)
}

type mysqlRepository struct {
}

func NewMysqlRepository() MysqlRepository {
	return &mysqlRepository{}
}

func (repository *mysqlRepository) Create(tx *sqlx.Tx, ctx context.Context, user modelentities.User) (rowsAffected int64, lastInsertedId int64, err error) {
	result, err := tx.ExecContext(ctx, `INSERT INTO users (email, password) VALUES (?, ?);`, user.Email, user.Password)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}
	lastInsertedId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func (repository *mysqlRepository) Get(db *sqlx.DB, ctx context.Context, id int) (user modelentities.User, err error) {
	err = db.GetContext(ctx, &user, `SELECT id, email, password FROM users WHERE id = ?;`, user.Id)
	return
}

func (repository *mysqlRepository) Update(tx *sqlx.Tx, ctx context.Context, user modelentities.User) (rowsAffected int64, err error) {
	result, err := tx.ExecContext(ctx, `UPDATE users SET email = ?, password = ? WHERE id = ?;`, user.Email, user.Password, user.Id)
	if err != nil {
		return
	}
	return result.RowsAffected()
}

func (repository *mysqlRepository) Delete(tx *sqlx.Tx, ctx context.Context, id int) (rowsAffected int64, err error) {
	result, err := tx.ExecContext(ctx, `DELETE FROM users WHERE id = ?;`, id)
	if err != nil {
		return
	}
	return result.RowsAffected()
}
