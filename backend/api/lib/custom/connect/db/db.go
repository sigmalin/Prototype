package db

import (
	"database/sql"
	"fmt"
	"log"

	"custom/config"

	_ "github.com/go-sql-driver/mysql"
)

var databases = make(map[string]*sql.DB)

func connect(tableName string) (*sql.DB, error) {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.SQL_USERNAME, config.SQL_PASSWORD, config.SQL_ADDRESS, config.SQL_PORT, tableName)
	db, err := sql.Open(config.SQL_DRIVER, conn)
	if err != nil {
		log.Fatal("connection to mysql failed : ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("connection to mysql failed : ", err)
	}

	db.SetMaxIdleConns(config.SQL_MAXIDLECONNECT)
	db.SetMaxOpenConns(config.SQL_MAXOPENCONNECT)

	return db, err
}

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
