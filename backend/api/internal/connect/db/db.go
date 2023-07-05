package db

import (
	"database/sql"
	"fmt"
	"log"

	"config"

	_ "github.com/go-sql-driver/mysql"
)

func connect(tableName string) (*sql.DB, error) {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.SQL_USERNAME, config.SQL_PASSWORD, config.SQL_ADDRESS, config.SQL_PORT, tableName)
	db, err := sql.Open(config.SQL_DRIVER, conn)
	if err != nil {
		log.Print("connection to mysql failed : ", err)
		return db, err
	}

	if err = db.Ping(); err != nil {
		log.Print("connection to mysql failed : ", err)
		return db, err
	}

	db.SetMaxIdleConns(config.SQL_MAXIDLECONNECT)
	db.SetMaxOpenConns(config.SQL_MAXOPENCONNECT)

	return db, err
}
