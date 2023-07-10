package allusers

import (
	"context"
	"database/sql"

	"response"
	"response/code"
)

type user struct {
	UserID int    `json:"UserID" example:"7"`
	Token  string `json:"Token" example:"d704e538-4f2f-486d-a2a1-a2b0ad3b4fe7"`
	Name   string `json:"Name" example:"sigma"`
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

	users := []user{}
	for rows.Next() {
		u := &user{}
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
