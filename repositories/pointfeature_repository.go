package repositories

import (
	"context"
	"go-rest-api/database"
	"go-rest-api/models"
	"log"
	"os"
)

func GetALlAccessPointFeatureswithFilter(filter interface{}) ([]models.PointFeature, error) {
	var features []models.PointFeature
	var pointFeature *models.PointFeature
	mongoManger := database.GetMongoManager()
	mongoClient := database.GetMongoManager().GetMongoAdapterClient()
	defer mongoManger.DisconnectMongoClient()

	collection := mongoClient.Database(os.Getenv("DB_DATABASE")).Collection(os.Getenv("MVP_ACCESS_COLLECTION"))
	ctx := context.TODO()
	cur, err := collection.Find(ctx, filter) //, options.Find().SetLimit(200)
	if err != nil {
		return features, err
	}
	for cur.Next(ctx) {
		pointFeature = &models.PointFeature{}
		er := cur.Decode(pointFeature)
		if er != nil {
			log.Println(er)
		}
		features = append(features, *pointFeature)
	}
	return features, err
}
