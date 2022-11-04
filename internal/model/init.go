package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

const MONGODB_URI = "mongodb://2019-nCov:123456@175.178.174.246:27017/?connect=direct"
const DATABASE_NAME = "2019-nCov"

func InitDataSource() {
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}
	DB = Client.Database(DATABASE_NAME)
}
