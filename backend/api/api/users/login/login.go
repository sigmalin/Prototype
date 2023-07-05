package login

import (
	"context"
	"database/sql"
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

func LogIn(args *Arguments) {

	prepare, err1 := args.db.PrepareContext(args.ctx, "SELECT UserID From Users WHERE Token = ?")
	if err1 != nil {
		args.response.Error(code.LOGIN_FAIURE, err1.Error())
		return
	}
	defer prepare.Close()

	var userID int64
	err2 := prepare.QueryRowContext(args.ctx, args.token).Scan(&userID)
	if err2 != nil {
		args.response.Error(code.LOGIN_FAIURE, err2.Error())
		return
	}

	err3 := updateLoginTime(args, userID)
	if err3 != nil {
		args.response.Error(code.LOGIN_FAIURE, err3.Error())
		return
	}

	args.session.Set(userID)

	// TODO : Login Data
	args.response.Message()
}

func updateLoginTime(args *Arguments, userID int64) error {
	prepare, err1 := args.db.PrepareContext(args.ctx, "Update Users SET UpdateTime = ? WHERE UserID = ?")
	if err1 != nil {
		return err1
	}
	defer prepare.Close()

	_, err2 := prepare.ExecContext(args.ctx, time.Now().Unix(), userID)
	return err2
}
