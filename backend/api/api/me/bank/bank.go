package bank

import (
	"context"
	"modal/user/bankData"
	"model/userData"
	"response"
	"response/code"

	"go.mongodb.org/mongo-driver/mongo"
)

type Arguments struct {
	db  *mongo.Database
	ctx context.Context
	id  string
}

type Result struct {
	Bank bankData.Content `json:"Bank"`
}

func NewArguments(db *mongo.Database, ctx context.Context, id string) *Arguments {
	return &Arguments{
		db:  db,
		ctx: ctx,
		id:  id,
	}
}

func Bank(args *Arguments, resp *response.Body) {

	data, err := userData.GetCache(args.ctx, args.db, args.id)

	if err != nil {
		resp.Error(code.DATA_NOT_FIND, err.Error())
		return
	}

	result := &Result{Bank: data.Bank}

	resp.Data = result
}
