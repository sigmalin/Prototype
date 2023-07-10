package signin

import (
	"context"
	"database/sql"
	"log"
	"response"
	"time"

	"response/code"
	"session"
)

type Arguments struct {
	db      *sql.DB
	ctx     context.Context
	session session.Session
	token   string
}

type signInData struct {
	Token string `json:"token"`
}

func NewArguments(db *sql.DB, ctx context.Context, session session.Session, token string) *Arguments {
	return &Arguments{
		db:      db,
		ctx:     ctx,
		session: session,
		token:   token,
	}
}

func SignIn(args *Arguments, res *response.Body) {

	prepare, err1 := args.db.PrepareContext(args.ctx, "INSERT INTO Users(Token, CreateTime, UpdateTime) VALUES(?, ?, ?)")
	if err1 != nil {
		log.Printf("createUser : %s", err1.Error())
		res.Error(code.SIGNIN_FAIURE, "createUser failure")
		return
	}
	defer prepare.Close()

	curTime := time.Now().Unix()
	row, err2 := prepare.ExecContext(args.ctx, args.token, curTime, curTime)
	if err2 != nil {
		log.Printf("createUser : %s", err2.Error())
		res.Error(code.SIGNIN_FAIURE, "createUser failure")
		return
	}

	userID, err3 := row.LastInsertId()
	if err3 != nil {
		log.Printf("createUser : %s", err3.Error())
		res.Error(code.SIGNIN_FAIURE, "createUser failure")
		return
	}

	args.session.Set(userID)

	res.Data = &signInData{Token: args.token}

}
