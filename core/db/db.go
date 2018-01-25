package db

import (
	"database/sql"
	"fmt"

	"github.com/syariatifaris/gosandbox/core/config"
	sqlx "github.com/syariatifaris/gosandbox/core/db/sqlx"
)

type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
	BeginTransaction() interface{}
	Execute(tx interface{}, query string, args ...interface{}) sql.Result
	CommitTransaction(tx interface{}) error
}

func NewRelationalDBConnection(dbConfig config.Database) (DB, error) {
	sqlx, err := sqlx.NewSqlxDbConnection(dbConfig)
	if err != nil {
		return nil, err
	}

	return sqlx, nil
}

func NewInjectRelationalDBConnection(cfg *config.ConfigurationData) DB {
	something := "Faris"
	fmt.Println(something)
	sqlxConn, _ := NewRelationalDBConnection(cfg.Database)
	return sqlxConn
}
