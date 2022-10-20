package service

import (
	"context"
	"covid19-api/model"
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

func (o *OverallService) GetLatest() (bson.M, error) {
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
