package bank

import (
	"context"
	"modal/user/bankData"
	"model/userData"
	"response"
	"response/code"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Arguments struct {
	db       *mongo.Database
	ctx      context.Context
	id       string
	coin     int64
	faith    int64
	gems     int64
	treasure int64
}

type Result struct {
	Bank bankData.Content `json:"Bank"`
}

func NewArguments(db *mongo.Database, ctx context.Context, id string, coin int64, faith int64, gems int64, treasure int64) *Arguments {
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

func Handle(args *Arguments, resp *response.Body) {

	userID, err := primitive.ObjectIDFromHex(args.id)
	if err != nil {
		resp.Error(code.INPUT_FAIURE, err.Error())
		return
	}

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "Bank", Value: bson.D{
			{Key: "Coin", Value: args.coin},
			{Key: "Faith", Value: args.faith},
			{Key: "Gems", Value: args.gems},
			{Key: "Treasure", Value: args.treasure},
		}},
	}}}

	result, err1 := args.db.Collection("users").UpdateOne(args.ctx, filter, update)
	if err1 != nil {
		resp.Error(code.DATA_NOT_FIND, err1.Error())
		return
	}

	if result.MatchedCount == 0 {
		resp.Error(code.DATA_NOT_FIND, "No match document")
		return
	}

	cache := userData.Exists(args.ctx, args.id)
	if cache != nil {
		cache.Bank.Coin = args.coin
		cache.Bank.Faith = args.faith
		cache.Bank.Gems = args.gems
		cache.Bank.Treasure = args.treasure
		userData.SetCache(args.id, cache)
	}

	resp.Data = &Result{
		Bank: bankData.Content{
			Coin:     args.coin,
			Faith:    args.faith,
			Gems:     args.gems,
			Treasure: args.treasure,
		},
	}
}
