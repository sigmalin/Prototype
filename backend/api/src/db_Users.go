package main

import (
	"context"
	"database/sql"
	"log"

	"custom/response"
	"custom/response/code"
)

/*
type User struct {
	Name string `db:"Name"`
	Mail string `db:"Mail"`
}
*/

type User struct {
	Name string `json:"Name"`
	Mail string `json:"Mail"`
}

/*
func queryAllUsers(db *sql.DB, res Response, ctx context.Context) string {
	rows, err := db.QueryContext(ctx, "SELECT * From Users")
	if err != nil {
		log.Print("[queryAllUsers] : ", err)
		res.Error(globalVar.UNKNOWN_ERROR, err.Error())
		return
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.Name, &u.Mail); err != nil {
			log.Print("[queryAllUsers] : ", err)
			res.Error(globalVar.UNKNOWN_ERROR, err.Error())
			return
		}
		users = append(users, *u)
	}

	err = rows.Err()
	if err != nil {
		log.Print("[queryAllUsers] : ", err)
		res.Error(globalVar.UNKNOWN_ERROR, err.Error())
		return
	}

	res.Add("data", users)
	res.Message()
}
*/

func queryUsers(db *sql.DB, res Response, ctx context.Context, name string) {
	rows, err := db.QueryContext(ctx, "SELECT * From Users Where Name = ?", name)
	if err != nil {
		log.Print("[queryAllUsers] : ", err)
		res.Error(code.UNKNOWN_ERROR, err.Error())
		return
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.Name, &u.Mail); err != nil {
			log.Print("[queryAllUsers] : ", err)
			res.Error(code.UNKNOWN_ERROR, err.Error())
			return
		}
		users = append(users, *u)
	}

	err = rows.Err()
	if err != nil {
		log.Print("[queryAllUsers] : ", err)
		res.Error(code.UNKNOWN_ERROR, err.Error())
		return
	}

	res.Add("data", users)
	res.Message()
}

func querySleep(db *sql.DB, res response.Response, ctx context.Context) {
	rows, err := db.QueryContext(ctx, "SELECT sleep(15)")
	if err != nil {
		log.Print("[queryAllUsers] : ", err)
		res.Error(code.UNKNOWN_ERROR, err.Error())
		return
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.Name, &u.Mail); err != nil {
			log.Print("[queryAllUsers] : ", err)
			res.Error(code.UNKNOWN_ERROR, err.Error())
			return
		}
		users = append(users, *u)
	}

	err = rows.Err()
	if err != nil {
		log.Print("[queryAllUsers] : ", err)
		res.Error(code.UNKNOWN_ERROR, err.Error())
	}

	res.Add("data", users)
	res.Message()
}
