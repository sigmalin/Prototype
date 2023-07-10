package allusers

import (
	"context"
	"database/sql"

	"response"
	"response/code"
)

type Response = response.Body

type User struct {
	UserID int    `json:"UserID"`
	Token  string `json:"Token"`
	Name   string `json:"Name"`
}

func Handle(ctx context.Context, db *sql.DB, res *response.Body) {

	prepare, err1 := db.PrepareContext(ctx, "SELECT UserID, Token, Name From Users")
	if err1 != nil {
		res.Error(code.UNKNOWN_ERROR, err1.Error())
		return
	}
	defer prepare.Close()

	rows, err2 := prepare.QueryContext(ctx)
	if err2 != nil {
		res.Error(code.UNKNOWN_ERROR, err2.Error())
		return
	}

	users := []User{}
	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.UserID, &u.Token, &u.Name); err != nil {
			res.Error(code.UNKNOWN_ERROR, err.Error())
			return
		}
		users = append(users, *u)
	}

	err3 := rows.Err()
	if err3 != nil {
		res.Error(code.UNKNOWN_ERROR, err3.Error())
		return
	}

	res.Data = users
}
