package infrastructure

import (
	"database/sql"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/aux-Issa/Next_Go_practice/app/interface/database"
)

type SqlHandler struct {
	Conn *sql.DB
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
	res := SqlHandler{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) database.Row {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlHandler), err
	}
	row := new(SqlHandler)
	row.Rows = rows
	return row, err
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
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
