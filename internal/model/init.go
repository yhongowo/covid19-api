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

const MONGODB_URI = ""
const DATABASE_NAME = "2019-nCov"

func InitDataSource() {
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic(err)
	}
	DB = Client.Database(DATABASE_NAME)
}
