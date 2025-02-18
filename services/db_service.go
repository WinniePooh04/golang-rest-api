package services

import (
	"go-rest-api/database"
	"log"
)

func InitDBAdapter() {
	database.GetDBAdapter()
	err := database.GetMongoManager().InitializeMongoAdapterClient()
	if err != nil {
		log.Println("Failed to connect to MongoDB:", err)
		panic(err) // This will stop execution if MongoDB connection fails
	}
	log.Println("Successfully connected to MongoDB")
}
