package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conn - sql connection handler
var Conn *sql.DB

// NewSQLHandler - init sql handler
func init() {
	user := "root"
	pass := "password"
	name := "itemsDB"

	dbconf := user + ":" + pass + "@/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error)
	}
	Conn = conn
}
