package service

import (
	"context"
	"covid19-api/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AbroadService struct {
	C          *gin.Context
	collection *mongo.Collection
}

func NewAbroadService(c *gin.Context) *AbroadService {
	return &AbroadService{C: c, collection: model.DB.Collection("Timeline")}
}

func (o *AbroadService) GetLatest() (result bson.M, err error) {
	if err = o.collection.FindOne(context.TODO(), bson.D{}).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}
