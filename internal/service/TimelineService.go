package service

import (
	"context"
	"covid19-api/internal/model"
	"covid19-api/internal/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TimelineService struct {
	C          *gin.Context
	collection *mongo.Collection
}

func NewTimelineService(c *gin.Context) *TimelineService {
	return &TimelineService{C: c, collection: model.DB.Collection("Timeline")}
}

//获取最新新闻数据
func (o *TimelineService) GetLatest() (result []bson.M, err error) {
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
