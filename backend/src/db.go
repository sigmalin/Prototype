package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MainDB *sql.DB

func dbConnect() {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", SQL_USERNAME, SQL_PASSWORD, SQL_ADDRESS, SQL_PORT, SQL_DATABASE)
	db, err := sql.Open(SQL_DRIVER, conn)
	if err != nil {
		log.Fatal("connection to mysql failed : ", err)
	}

	for i := 0; i < SQL_RETRYMAX; i++ {
		if err = db.Ping(); err != nil {
			time.Sleep(SQL_RETRYINTERVAL)
			log.Printf("connet db failure, retry : %d", i)
		} else {
			break
		}
	}

	if err != nil {
		log.Fatal("connection to mysql failed : ", err)
	}

	db.SetMaxIdleConns(SQL_MAXIDLECONNECT)
	db.SetMaxOpenConns(SQL_MAXOPENCONNECT)

	MainDB = db
}
