package service

import (
	"context"
	"covid19-api/internal/model"
	"covid19-api/internal/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AreaService struct {
	C          *gin.Context
	collection *mongo.Collection
}

func NewAreaService(c *gin.Context) *AreaService {
	return &AreaService{C: c, collection: model.DB.Collection("Area")}
}

//获取最新所有省份，城市的数据
func (o *AreaService) ListProvince() ([]bson.M, error) {
	filter := bson.D{
		{"updateTime", bson.D{{"$gte", util.TodayBeginTime()}}},
	}
	cursor, err := o.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	results := make([]bson.M, 0)
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
