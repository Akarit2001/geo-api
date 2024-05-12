package databases

import (
	"context"
	"errors"
	"geo-test/modules/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(cfg *configs.Config, ctx context.Context) (*mongo.Database, error) {

	if cfg.MongoConnection == "" || cfg.MongoDBName == "" {
		return nil, errors.New("URI is empty")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoConnection))

	if err != nil {
		return nil, err
	}

	return client.Database(cfg.MongoDBName), nil
}
