package signin

import (
	"context"
	"modal/user/bankData"
	"response"
	"response/code"
	"time"

	"jwtGin"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
type jsonWebToken = *tokenData.Content

type Result struct {
	AccessToken string `json:"AccessToken,omitempty" example:"d704e538-4f2f-486d-a2a1-a2b0ad3b4fe7"`
	jsonWebToken
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

func SignIn(args *Arguments, resp *response.Body) {
	id, err1 := createAccount(args)
	if err1 != nil {
		resp.Error(code.SIGNIN_FAIURE, err1.Error())
		return
	}

	err2 := createUser(args, id)
	if err2 != nil {
		resp.Error(code.SIGNIN_FAIURE, err2.Error())
		return
	}

	userID := id.Hex()

	info, err3 := getUserInfo(args, userID)
	if err3 != nil {
		resp.Error(code.SIGNIN_FAIURE, err3.Error())
		return
	}

	data, err4 := getResult(args, userID, info)
	if err4 != nil {
		resp.Error(code.SIGNIN_FAIURE, err4.Error())
		return
	}

	resp.Data = data
}

func createAccount(args *Arguments) (primitive.ObjectID, error) {

	data := &accountData.Content{
		Token:      args.token,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	result, err := args.db.Collection("accounts").InsertOne(args.ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil

}

func createUser(args *Arguments, id primitive.ObjectID) error {

	bankData, err1 := bankData.NewContent()
	if err1 != nil {
		return err1
	}

	data := &userData.Content{
		ID:   id,
		Bank: *bankData,
	}

	_, err2 := args.db.Collection("users").InsertOne(args.ctx, data)
	if err2 != nil {
		return err2
	}

	return nil
}

func getUserInfo(args *Arguments, id string) (*userData.Content, error) {
	return userData.GetCache(args.ctx, args.db, id)
}

func getResult(args *Arguments, id string, info *userData.Content) (*Result, error) {
	token, err := tokenData.NewContent(id, args.jwtMgr)
	if err != nil {
		return nil, err
	}

	return &Result{
		AccessToken:  args.token,
		jsonWebToken: token,
		userInfo:     info,
	}, nil
}
