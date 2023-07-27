package userData

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"modal/user/bankData"

	cacher "cache/redis"
)

type Content struct {
	ID   primitive.ObjectID `json:"-" bson:"_id"`
	Bank bankData.Content   `json:"Bank" bson:"Bank"`
}

func cacheKey(UserID string) string {
	return "user/" + UserID
}

func NewContent(ctx context.Context, db *mongo.Database, id string) (interface{}, error) {
	userID, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		return nil, err1
	}

	filter := bson.D{{Key: "_id", Value: userID}} //bson.M{"_id": userID}

	var data Content
	err2 := db.Collection("users").FindOne(ctx, filter).Decode(&data)
	if err2 != nil {
		return nil, err2
	}

	return &data, nil
}

func GetCache(ctx context.Context, db *mongo.Database, id string) (*Content, error) {
	value := new(Content)
	err := cacher.Search(cacheKey(id), value, func() (interface{}, error) {
		return NewContent(ctx, db, id)
	})
	if err != nil {
		return nil, err
	}
	return value, err
}

func SetCache(id string, value *Content) error {
	return cacher.Set(cacheKey(id), value)
}

func DelCache(ctx context.Context, id string) error {
	return cacher.Delete(ctx, cacheKey(id))
}

func Exists(ctx context.Context, id string) *Content {
	key := cacheKey(id)
	exist := cacher.Exists(ctx, key)
	if !exist {
		return nil
	}
	var value Content
	err := cacher.Get(ctx, key, &value)
	if err != nil {
		return nil
	}
	return &value
}
