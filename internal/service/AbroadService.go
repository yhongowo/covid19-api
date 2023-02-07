package service

import (
	"context"
	"covid19-api/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AbroadService struct {
	C          *gin.Context
	collection *mongo.Collection
}

func NewAbroadService(c *gin.Context) *AbroadService {
	return &AbroadService{C: c, collection: model.DB.Collection("Abroad")}
}

func (o *AbroadService) GetLatest() (result []bson.M, err error) {
	//查询最近时间戳
	t := &updateTime{}
	opts := options.FindOne().SetProjection(bson.D{{"updateTime", 1}}).SetSort(bson.D{{"updateTime", -1}})
	if err := o.collection.FindOne(context.TODO(), bson.D{}, opts).Decode(t); err != nil {
		return nil, err
	}

	//查询最新数据
	cursor, err := o.collection.Find(context.TODO(), bson.D{{"updateTime", t.UpdateTime}})
	if err != nil {
		return nil, err
	}
	results := make([]bson.M, 0)
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
