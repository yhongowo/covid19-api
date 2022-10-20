package model

import (
	"context"
	"covid19-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

func InitDataSource() {
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGODB_URI))
	if err != nil {
		panic(err)
	}
	DB = Client.Database(config.DATABASE_NAME)
}
