package bank

import (
	"context"
	"database/sql"
	"modal/bankData"
	"response"
	"response/code"
)

type Arguments struct {
	db  *sql.DB
	ctx context.Context
	id  string
}

func NewArguments(db *sql.DB, ctx context.Context, id string) *Arguments {
	return &Arguments{
		db:  db,
		ctx: ctx,
		id:  id,
	}
}

func Bank(args *Arguments, res *response.Body) {

	prepare, err1 := args.db.PrepareContext(args.ctx, "SELECT Coin, Faith, Gems, Treasure From Bank WHERE UserID = ?")
	if err1 != nil {
		res.Error(code.DATA_NOT_FIND, err1.Error())
		return
	}
	defer prepare.Close()

	data, _ := bankData.NewContent()

	err2 := prepare.QueryRowContext(args.ctx, args.id).Scan(&data.Coin, &data.Faith, &data.Gems, &data.Treasure)
	if err2 != nil {
		res.Error(code.UNKNOWN_ERROR, err2.Error())
		return
	}

	res.Data = data
}
