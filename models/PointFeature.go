package models

import "time"

type PointFeature struct {
	Type       string          `json:"type" bson:"type"`
	Geometry   PointGeometry   `json:"geometry" bson:"geometry"`
	Properties PointProperties `json:"properties" bson:"properties"`
}
type PointGeometry struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
type PointProperties struct {
	Type           string            `json:"type" bson:"type"`
	Id             string            `json:"id" bson:"id"`
	AdditionalInfo map[string]string `json:"additional_info" bson:"additional_info"`
	Tag            []string          `json:"tags" bson:"tags"`
	SourceChannel  string            `json:"-" bson:"source_channel"`
	Created_on     time.Time         `json:"created_on"   bson:"created_on"`
	Updated_on     time.Time         `json:"updated_on"   bson:"updated_on"`
}
