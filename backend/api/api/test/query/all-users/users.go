package allusers

import (
	"context"
	"model/accountData"

	"response"
	"response/code"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Arguments struct {
	db    *mongo.Database
	ctx   context.Context
	start int64
	count int64
}

type Result = accountData.Content

func NewArguments(db *mongo.Database, ctx context.Context, start int64, count int64) *Arguments {
	return &Arguments{
		db:    db,
		ctx:   ctx,
		start: start,
		count: count,
	}
}

func Handle(args *Arguments, resp *response.Body) {

	opts := options.Find()
	opts.SetSkip(args.start)
	opts.SetLimit(args.count)

	cursor, err1 := args.db.Collection("accounts").Find(args.ctx, bson.D{{}}, opts)
	if err1 != nil {
		resp.Error(code.UNKNOWN_ERROR, err1.Error())
		return
	}
	defer func() {
		err := cursor.Close(context.Background())
		if err != nil {
			resp.Error(code.UNKNOWN_ERROR, err.Error())
			return
		}
	}()

	var results []Result
	err2 := cursor.All(args.ctx, &results)
	if err2 != nil {
		resp.Error(code.UNKNOWN_ERROR, err2.Error())
		return
	}

	resp.Data = results
}
