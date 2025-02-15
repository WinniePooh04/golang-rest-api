package services

import (
	"go-rest-api/database"
	"log"
)

func InitDBAdapter() {
	database.GetDBAdapter()
	err := database.GetMongoManager().InitializeMongoAdapterClient()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB:")
	} else {
		log.Fatalf("Success to connect MongoDB:")
	}
}
