package signin

import (
	"context"
	"database/sql"
	"response"
	"time"

	"jwtGin"
	"model/signinData"
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

func SignIn(args *Arguments, res *response.Body) {

	userID, err := createUser(args)
	if err != nil {
		res.Error(code.SIGNIN_FAIURE, "createUser failure : "+err.Error())
		return
	}

	err1 := createBank(args, userID)
	if err1 != nil {
		res.Error(code.SIGNIN_FAIURE, "createBank failure : "+err1.Error())
		return
	}

	data, err2 := signinData.NewContent(args.token, userID, args.jwtMgr)
	if err2 != nil {
		res.Error(code.SIGNIN_FAIURE, err2.Error())
		return
	}

	res.Data = data
}

func createUser(args *Arguments) (int64, error) {

	prepare, err1 := args.db.PrepareContext(args.ctx, "INSERT INTO Users(Token, CreateTime, UpdateTime) VALUES(?, ?, ?)")
	if err1 != nil {
		return 0, err1
	}
	defer prepare.Close()

	curTime := time.Now().Unix()
	row, err2 := prepare.ExecContext(args.ctx, args.token, curTime, curTime)
	if err2 != nil {
		return 0, err2
	}

	userID, err3 := row.LastInsertId()
	if err3 != nil {
		return 0, err3
	}

	return userID, nil
}

func createBank(args *Arguments, userID int64) error {

	prepare, err1 := args.db.PrepareContext(args.ctx, "INSERT INTO Bank(UserID, Coin, Faith, Gems, Treasure) VALUES(?, ?, ?, ?, ?)")
	if err1 != nil {
		return err1
	}
	defer prepare.Close()

	_, err2 := prepare.ExecContext(args.ctx, userID, 1000, 0, 0, 0)
	if err2 != nil {
		return err2
	}

	return nil
}
