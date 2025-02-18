package services

import (
	"go-rest-api/models"
	"go-rest-api/repositories"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPointFeatureswithFilter(body models.FTTxRequestBody) ([]models.PointFeature, error) {
	mainFilter := bson.M{}
	filters := bson.A{}
	deviceType := strings.Split(body.Type, ",")
	queryFilters := bson.A{
		bson.M{"geometry.coordinates.0": bson.M{"$gte": body.SWLat, "$lte": body.NELat}},
		bson.M{"geometry.coordinates.1": bson.M{"$gte": body.SWLng, "$lte": body.NELng}},
		bson.M{"properties.type": deviceType[0]},
		bson.M{"properties.additional_info.display": "true"},
	}
	filter := bson.M{"$and": queryFilters}
	filters = append(filters, filter)
	mainFilter = bson.M{"$or": filters}
	return repositories.GetALlAccessPointFeatureswithFilter(mainFilter)
}
