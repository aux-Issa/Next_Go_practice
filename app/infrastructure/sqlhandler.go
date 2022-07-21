package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(conn *sql.DB) *SqlHandler {
	conn, err := sql.Open(driverName: "mysql", dataSourceName: "")
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}