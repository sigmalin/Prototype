package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

func checkUserExist(db *sql.DB, ctx context.Context, mail string) bool {

	prepare, err1 := db.PrepareContext(ctx, "SELECT UserID From Users WHERE Mail = ?")
	if err1 != nil {
		log.Print("checkUserExist : sql prepare failure")
		return false
	}
	defer prepare.Close()

	row := prepare.QueryRowContext(ctx, mail)

	var userID int64
	err2 := row.Scan(&userID)
	return err2 == nil
}

func createUser(db *sql.DB, ctx context.Context, name string, mail string, pw string) (int64, error) {

	prepare, err1 := db.PrepareContext(ctx, "INSERT INTO Users(Name, Mail, Password, CreateTime, UpdateTime) VALUES(?, ?, ?, ?, ?)")
	if err1 != nil {
		log.Printf("createUser : %s", err1.Error())
		return -1, errors.New("createUser failure")
	}
	defer prepare.Close()

	curTime := time.Now().Unix()
	res, err2 := prepare.ExecContext(ctx, name, mail, pw, curTime, curTime)
	if err2 != nil {
		log.Printf("createUser : %s", err2.Error())
		return -1, errors.New("createUser failure")
	}

	userID, err3 := res.LastInsertId()
	if err3 != nil {
		log.Printf("createUser : %s", err3.Error())
		return -1, errors.New("createUser failure")
	}
	return userID, err3
}

func ExecSignin(db *sql.DB, ctx context.Context, name string, mail string, pw string) (int64, error) {

	if exist := checkUserExist(db, ctx, mail); exist {
		return -1, errors.New("mail has exist")
	}

	userID, err := createUser(db, ctx, name, mail, pw)
	return userID, err
}
