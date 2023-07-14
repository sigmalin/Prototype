package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var databases = make(map[string]*sql.DB)

func GetDB(tableName string) *sql.DB {

	db, ok := databases[tableName]
	if !ok {
		conn, err := connect(tableName)
		if err != nil {
			return nil
		}
		databases[tableName] = conn
		db = conn
	}
	return db
}
