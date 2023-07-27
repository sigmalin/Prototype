package login

import (
	"context"
	"response"
	"response/code"
	"time"

	"jwtGin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"model/accountData"
	"model/user/tokenData"
	"model/userData"
)

type Arguments struct {
	db     *mongo.Database
	ctx    context.Context
	token  string
	jwtMgr *jwtGin.Manager
}

type userInfo = *userData.Content
type token = *tokenData.Content

type Result struct {
	token
	userInfo
}

func NewArguments(db *mongo.Database, ctx context.Context, token string, jwtMgr *jwtGin.Manager) *Arguments {
	return &Arguments{
		db:     db,
		ctx:    ctx,
		token:  token,
		jwtMgr: jwtMgr,
	}
}

func LogIn(args *Arguments, resp *response.Body) {
	id, err1 := auth(args)
	if err1 != nil {
		resp.Error(code.AUTH_FAIURE, err1.Error())
		return
	}

	info, err2 := getUserInfo(args, id)
	if err2 != nil {
		resp.Error(code.AUTH_FAIURE, err2.Error())
		return
	}

	data, err3 := getResult(args, id, info)
	if err3 != nil {
		resp.Error(code.AUTH_FAIURE, err3.Error())
		return
	}

	resp.Data = data
}

func auth(args *Arguments) (string, error) {
	filter := bson.D{{Key: "Token", Value: args.token}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "UpdateTime", Value: time.Now().Unix()},
	}}}

	collect := args.db.Collection("accounts")

	var account accountData.ID

	err1 := collect.FindOne(args.ctx, filter).Decode(&account)
	if err1 != nil {
		return "", err1
	}

	_, err2 := collect.UpdateOne(args.ctx, filter, update)
	if err2 != nil {
		return "", err2
	}

	return account.ID.Hex(), nil
}

func getUserInfo(args *Arguments, id string) (*userData.Content, error) {
	return userData.GetCache(args.ctx, args.db, id)
}

func getResult(args *Arguments, id string, info *userData.Content) (*Result, error) {
	data, err := tokenData.NewContent(id, args.jwtMgr)
	if err != nil {
		return nil, err
	}

	return &Result{token: data, userInfo: info}, nil
}
