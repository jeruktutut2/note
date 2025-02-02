package utils

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresUtil interface {
	GetDb() *sqlx.DB
	BeginTxx(ctx context.Context, options *sql.TxOptions) (*sqlx.Tx, error)
	Close(host string, port string)
	CommitOrRollback(tx *sqlx.Tx, err error) error
}

type PostgresUtilImplementation struct {
	db *sqlx.DB
}

func NewPostgresUtil(host string, username string, password string, database string, port string, applicationName string, maxOpenConnection int, maxIdleConnection int, connectionMaxIdletime int, connectionMaxLifetime int) PostgresUtil {
	println(time.Now().String(), "postgres: connecting to", host, ":", port)
	db, err := sqlx.Connect("postgres", "user="+username+" dbname="+database+" sslmode=disable password="+password+" host="+host+" port="+port+" application_name="+applicationName)
	if err != nil {
		log.Fatalln("postgres: error when connecting to:", err)
	}

	// maxOpenConnectionEnv := os.Getenv("POSTGRES_MAX_OPEN_CONNECTION")
	// maxOpenConnection, err := strconv.Atoi(maxOpenConnectionEnv)
	// if err != nil {
	// 	log.Fatalln("postgres: error when converting max open connection from string to int:", err)
	// }
	db.SetMaxOpenConns(maxOpenConnection)

	// maxIdleConnectionEnv := os.Getenv("POSTGRES_MAX_IDLE_CONNECTION")
	// maxIdleConnection, err := strconv.Atoi(maxIdleConnectionEnv)
	// if err != nil {
	// 	log.Fatalln("postgres: error when converting max idle connection from string to int", err)
	// }
	db.SetMaxIdleConns(maxIdleConnection)

	// connectionMaxIdletimeEnv := os.Getenv("POSTGRES_CONNECTION_MAX_IDLETIME")
	// connectionMaxIdletime, err := strconv.Atoi(connectionMaxIdletimeEnv)
	// if err != nil {
	// 	log.Fatalln("postgres: error when converting connection max idletime from string to int:", err)
	// }
	db.SetConnMaxLifetime(time.Duration(connectionMaxIdletime) * time.Minute)

	// connectionMaxLifetimeEnv := os.Getenv("POSTGRES_CONNECTION_MAX_LIFETIME")
	// connectionMaxLifetime, err := strconv.Atoi(connectionMaxLifetimeEnv)
	// if err != nil {
	// 	log.Fatalln("postgres: error when converting connection max lifetime from string to int:", err)
	// }
	db.SetConnMaxLifetime(time.Duration(connectionMaxLifetime) * time.Minute)
	println(time.Now().String(), "postgres: connected to", host, ":", port)

	println(time.Now().String(), "postgres: pinging to", host, ":", port)
	err = db.Ping()
	if err != nil {
		log.Fatalln("postgres: error when pinging to:", err)
	}
	println(time.Now().String(), "postgres: pinged to", host, ":", port)

	// println(time.Now().String(), "postgres: connected to", os.Getenv("POSTGRES_HOST"), ":", os.Getenv("POSTGRES_PORT"))
	return &PostgresUtilImplementation{
		db: db,
	}
}

func (util *PostgresUtilImplementation) GetDb() *sqlx.DB {
	return util.db
}

func (util *PostgresUtilImplementation) BeginTxx(ctx context.Context, options *sql.TxOptions) (*sqlx.Tx, error) {
	return util.db.BeginTxx(ctx, options)
}

func (util *PostgresUtilImplementation) Close(host string, port string) {
	println(time.Now().String(), "postgres: closing to", host, ":", port)
	err := util.db.Close()
	if err != nil {
		log.Fatalln("postgres: error when closing to:", err)
	}
	println(time.Now().String(), "postgres: closed to", host, ":", port)
}

func (util *PostgresUtilImplementation) CommitOrRollback(tx *sqlx.Tx, err error) error {
	if err == nil {
		err = tx.Commit()
		if err != nil && err != sql.ErrTxDone {
			err = tx.Rollback()
			if err != nil && err != sql.ErrTxDone {
				return err
			}
			return nil
		}
		return nil
	} else {
		err = tx.Rollback()
		if err != nil && err != sql.ErrTxDone {
			return err
		}
		return nil
	}
}
