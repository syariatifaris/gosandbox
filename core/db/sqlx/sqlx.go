package sqlx

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/gosandbox/core/config"
)

const DSCFormat = `%s:%s@tcp(%s:%d)/%s`

type Sqlx struct {
	db *sqlx.DB
}

func NewSqlxDbConnection(dbConfig config.Database) (*Sqlx, error) {
	//"<username>:<pw>@tcp(<HOST>:<port>)/<dbname>"
	connectionString := fmt.Sprintf(DSCFormat, dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
	)

	db, err := sqlx.Connect(dbConfig.DBType, connectionString)

	if err != nil {
		return nil, err
	}

	return &Sqlx{
		db: db,
	}, nil
}

func (s *Sqlx) Select(dest interface{}, query string, args ...interface{}) error {
	if len(args) > 0 {
		return s.db.Select(dest, query, args...)
	}
	return s.db.Select(dest, query)
}

func (s *Sqlx) BeginTransaction() interface{} {
	return s.db.MustBegin()
}

func (s *Sqlx) Execute(tx interface{}, query string, args ...interface{}) sql.Result {
	if _, ok := tx.(*sqlx.Tx); !ok {
		return nil
	}
	return tx.(*sqlx.Tx).MustExec(query, args...)
}

func (s *Sqlx) CommitTransaction(tx interface{}) error {
	if _, ok := tx.(*sqlx.Tx); !ok {
		return fmt.Errorf("wrong argument type, sqlx.DB needed")
	}
	return tx.(*sqlx.Tx).Commit()
}
