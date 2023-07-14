package login

import (
	"context"
	"database/sql"
	"response"
	"time"

	"jwtGin"
	"model/loginData"
	"response/code"
)

type Arguments struct {
	db     *sql.DB
	ctx    context.Context
	token  string
	jwtMgr *jwtGin.Manager
}

func NewArguments(db *sql.DB, ctx context.Context, token string, jwtMgr *jwtGin.Manager) *Arguments {
	return &Arguments{
		db:     db,
		ctx:    ctx,
		token:  token,
		jwtMgr: jwtMgr,
	}
}

func LogIn(args *Arguments, res *response.Body) {

	prepare, err1 := args.db.PrepareContext(args.ctx, "SELECT UserID From Users WHERE Token = ?")
	if err1 != nil {
		res.Error(code.LOGIN_FAIURE, err1.Error())
		return
	}
	defer prepare.Close()

	var userID int64
	err2 := prepare.QueryRowContext(args.ctx, args.token).Scan(&userID)
	if err2 != nil {
		res.Error(code.LOGIN_FAIURE, err2.Error())
		return
	}

	err3 := updateLoginTime(args, userID)
	if err3 != nil {
		res.Error(code.LOGIN_FAIURE, err3.Error())
		return
	}

	data, err4 := loginData.NewContent(userID, args.jwtMgr)
	if err4 != nil {
		res.Error(code.LOGIN_FAIURE, err4.Error())
		return
	}

	res.Data = data
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
