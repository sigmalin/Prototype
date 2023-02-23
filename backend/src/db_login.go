package main

import (
	"context"
	"database/sql"
	"time"
)

func QueryLogin(db *sql.DB, ctx context.Context, mail string, pw string) (int64, error) {

	prepare, err1 := db.PrepareContext(ctx, "SELECT UserID From Users WHERE Mail = ? AND Password = ?")
	if err1 != nil {
		return -1, err1
	}
	defer prepare.Close()

	var userID int64
	err2 := prepare.QueryRowContext(ctx, mail, pw).Scan(&userID)
	return userID, err2
}

func UpdateLoginTime(db *sql.DB, ctx context.Context, userID int64) error {

	prepare, err1 := db.PrepareContext(ctx, "Update Users SET UpdateTime = ? WHERE UserID = ?")
	if err1 != nil {
		return err1
	}
	defer prepare.Close()

	_, err2 := prepare.ExecContext(ctx, time.Now().Unix(), userID)
	return err2
}
