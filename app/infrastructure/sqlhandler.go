package infrastructure

import (
	"database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/aux-Issa/Next_Go_practice/app/interface/database"
)

type SqlHandler struct {
	Conn *sql.DB
}
type SqlResult struct {
	Result sql.Result
}
type SqlRow struct {
	Rows *sql.Rows
}

func NewSqlHandler(conn *sql.DB) *SqlHandler {
	conn, err := sql.Open("mysql", "")
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := new(SqlResult)
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return nil, err
	}
	row := new(SqlRow)
	row.Rows = rows
	// ⇩ここrowかもしれない
	return rows, err
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (r SqlRow) Scan(...interface{}) error {
	return r.Rows.Scan()
}

func (r SqlRow) Next(...interface{}) bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
