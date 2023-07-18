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

	data, err := bankData.GetCache(args.ctx, args.db, args.id)

	if err != nil {
		res.Error(code.DATA_NOT_FIND, err.Error())
		return
	}

	res.Data = data
}
