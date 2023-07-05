package signin

import (
	"context"
	"database/sql"
	"log"
	"time"

	"response"
	"response/code"
	"session"
)

type Arguments struct {
	db       *sql.DB
	ctx      context.Context
	response response.Response
	session  session.Session
	token    string
}

func NewArguments(db *sql.DB, ctx context.Context, response response.Response, session session.Session, token string) *Arguments {
	return &Arguments{
		db:       db,
		ctx:      ctx,
		response: response,
		session:  session,
		token:    token,
	}
}

func SignIn(args *Arguments) {

	prepare, err1 := args.db.PrepareContext(args.ctx, "INSERT INTO Users(Token, CreateTime, UpdateTime) VALUES(?, ?, ?)")
	if err1 != nil {
		log.Printf("createUser : %s", err1.Error())
		args.response.Error(code.SIGNIN_FAIURE, "createUser failure")
		return
	}
	defer prepare.Close()

	curTime := time.Now().Unix()
	res, err2 := prepare.ExecContext(args.ctx, args.token, curTime, curTime)
	if err2 != nil {
		log.Printf("createUser : %s", err2.Error())
		args.response.Error(code.SIGNIN_FAIURE, "createUser failure")
		return
	}

	userID, err3 := res.LastInsertId()
	if err3 != nil {
		log.Printf("createUser : %s", err3.Error())
		args.response.Error(code.SIGNIN_FAIURE, "createUser failure")
		return
	}

	args.session.Set(userID)

	args.response.Add("token", args.token)

	args.response.Message()

}
