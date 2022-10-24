package service

import (
	"context"
	"covid19-api/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type OverallService struct {
	C          *gin.Context
	collection *mongo.Collection
}

func NewOverallService(c *gin.Context) *OverallService {
	return &OverallService{C: c, collection: model.DB.Collection("Overall")}
}

//获取最新overall信息
func (o *OverallService) GetOverall() (bson.M, error) {
	opts := options.Find().SetSort(bson.D{{"id", -1}}).SetLimit(1)
	cursor, err := o.collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results[0], nil
}

func (o *OverallService) ListOverall() ([]bson.M, error) {
	cursor, err := o.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, nil
}
