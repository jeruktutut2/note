package repositories

import (
	"context"

	modelentities "note-golang-postgresql/models/entities"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository interface {
	Create(tx *sqlx.Tx, ctx context.Context, user modelentities.User) (id int, err error)
	Get(db *sqlx.DB, ctx context.Context, id int) (user modelentities.User, err error)
	Update(tx *sqlx.Tx, ctx context.Context, user modelentities.User) (rowsAffected int64, err error)
	Delete(tx *sqlx.Tx, ctx context.Context, id int) (rowsAffected int64, err error)
}

type postgresRepository struct {
}

func NewPostgresRepository() PostgresRepository {
	return &postgresRepository{}
}

func (repository *postgresRepository) Create(tx *sqlx.Tx, ctx context.Context, user modelentities.User) (id int, err error) {
	err = tx.GetContext(ctx, &id, `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id;`, user.Email, user.Password)
	return
}

func (repository *postgresRepository) Get(db *sqlx.DB, ctx context.Context, id int) (user modelentities.User, err error) {
	err = db.GetContext(ctx, &user, `SELECT id, email, password FROM users WHERE id = $1;`, user.Id)
	return
}

func (repository *postgresRepository) Update(tx *sqlx.Tx, ctx context.Context, user modelentities.User) (rowsAffected int64, err error) {
	result, err := tx.ExecContext(ctx, `UPDATE users SET email = $1, password = $2 WHERE id = $3;`, user.Email, user.Password, user.Id)
	if err != nil {
		return
	}
	return result.RowsAffected()
}

func (repository *postgresRepository) Delete(tx *sqlx.Tx, ctx context.Context, id int) (rowsAffected int64, err error) {
	result, err := tx.ExecContext(ctx, `DELETE FROM users WHERE id = $1;`, id)
	if err != nil {
		return
	}
	return result.RowsAffected()
}
