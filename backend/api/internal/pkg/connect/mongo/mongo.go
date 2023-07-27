package cmongo

import (
	"context"
	"fmt"
	"log"

	"config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(ctx context.Context) (*mongo.Client, error) {

	credential := options.Credential{
		Username: config.DATABASE_USERNAME,
		Password: config.DATABASE_PASSWORD,
	}

	conn := fmt.Sprintf("mongodb://%s:%d", config.DATABASE_ADDRESS, config.DATABASE_PORTS)

	opts := options.Client().
		ApplyURI(conn).
		SetAuth(credential)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Print("connection to mongoDB failed : ", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Print("connection to mongoDB failed : ", err)
		return nil, err
	}

	return client, nil
}
