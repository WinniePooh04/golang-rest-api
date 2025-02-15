package services

import (
	"go-rest-api/models"
	"go-rest-api/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPointFeatureswithFilter(body models.FTTxRequestBody) ([]models.PointFeature, error) {
	filter := bson.M{}
	return repositories.GetALlAccessPointFeatureswithFilter(filter)
}
