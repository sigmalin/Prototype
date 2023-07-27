package cmongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

var databases = make(map[string]*mongo.Database)

func NewClient(ctx context.Context) error {
	if client == nil {
		var err error
		client, err = connect(ctx)
		return err
	}
	return nil
}

func GetDB(tableName string) *mongo.Database {

	if client == nil {
		log.Fatal("DB has not initialize")
		return nil
	}

	db, ok := databases[tableName]
	if !ok {
		db = client.Database(tableName)
		databases[tableName] = db
	}
	return db
}

func Release(ctx context.Context) {

	if client != nil {
		client.Disconnect(ctx)
		client = nil
	}

	databases = make(map[string]*mongo.Database)
}
