package service

import (
	"context"
	"covid19-api/internal/model"
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

//get latest data
func (o *TimelineService) GetLatest() (result bson.M, err error) {
	if err = o.collection.FindOne(context.TODO(), bson.D{}).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}
