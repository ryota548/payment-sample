package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Conn - sql connection handler
var Conn *sql.DB

// NewSQLHandler - init sql handler
func init() {
	user := os.Getenv("root")
	pass := os.Getenv("root")
	name := os.Getenv("itemsDB")

	dbconf := user + ":" + pass + "@/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error)
	}
	Conn = conn
}
