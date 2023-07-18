package bank

import (
	"context"
	"database/sql"
	"modal/bankData"
	"response"
	"response/code"
)

type Arguments struct {
	db       *sql.DB
	ctx      context.Context
	id       string
	coin     int64
	faith    int64
	gems     int64
	treasure int64
}

func NewArguments(db *sql.DB, ctx context.Context, id string, coin int64, faith int64, gems int64, treasure int64) *Arguments {
	return &Arguments{
		db:       db,
		ctx:      ctx,
		id:       id,
		coin:     coin,
		faith:    faith,
		gems:     gems,
		treasure: treasure,
	}
}

func Handle(args *Arguments, res *response.Body) {

	prepare, err1 := args.db.PrepareContext(args.ctx, "Update Bank SET Coin = ?, Faith = ?, Gems = ?, Treasure = ? WHERE UserID = ?")
	if err1 != nil {
		res.Error(code.DATA_NOT_FIND, err1.Error())
		return
	}
	defer prepare.Close()

	_, err2 := prepare.ExecContext(args.ctx, args.coin, args.faith, args.gems, args.treasure, args.id)
	if err2 != nil {
		res.Error(code.UNKNOWN_ERROR, err2.Error())
		return
	}

	err3 := bankData.DelCache(args.ctx, args.id)
	if err3 != nil {
		res.Error(code.UNKNOWN_ERROR, err3.Error())
		return
	}

	data, err4 := bankData.GetCache(args.ctx, args.db, args.id)
	if err4 != nil {
		res.Error(code.UNKNOWN_ERROR, err4.Error())
		return
	}

	res.Data = data
}
