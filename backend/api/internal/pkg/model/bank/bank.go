package bankData

import (
	"context"
	"database/sql"
	"errors"

	cacher "cache/redis"
)

type Content struct {
	Coin     int64 `json:"Coin" example:"1000"`
	Faith    int64 `json:"Faith" example:"0"`
	Gems     int64 `json:"Gems" example:"0"`
	Treasure int64 `json:"Treasure" example:"0"`
}

func CacheKey(UserID string) string {
	return "me/bank/" + UserID
}

func NewContent(ctx context.Context, db *sql.DB, id string) (interface{}, error) {

	prepare, err1 := db.PrepareContext(ctx, "SELECT Coin, Faith, Gems, Treasure From Bank WHERE UserID = ?")
	if err1 != nil {
		return nil, err1
	}
	defer prepare.Close()

	data := &Content{}

	err2 := prepare.QueryRowContext(ctx, id).Scan(&data.Coin, &data.Faith, &data.Gems, &data.Treasure)
	if err2 != nil {
		return nil, err2
	}

	return data, nil
}

func GetCache(ctx context.Context, db *sql.DB, id string) (interface{}, error) {
	return cacher.Search(CacheKey(id), func() (interface{}, error) {
		return NewContent(ctx, db, id)
	})
}

func SetCache(id string, content *Content) (interface{}, error) {
	return cacher.Search(CacheKey(id), func() (interface{}, error) {
		if content == nil {
			return nil, errors.New("content is nil")
		}
		return content, nil
	})
}

func DelCache(ctx context.Context, id string) error {
	return cacher.Delete(ctx, CacheKey(id))
}
