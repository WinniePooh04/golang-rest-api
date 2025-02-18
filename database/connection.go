package database

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/banyar/go-packages/pkg/adapters"
	"github.com/banyar/go-packages/pkg/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoManager struct {
	client *mongo.Client
}

var (
	instance     *mongoManager
	once         sync.Once
	mongoAdapter *adapters.MongoAdapter
)

type DSNMongo struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

// GetMongoManager returns the singleton instance of mongoManager
func GetMongoManager() *mongoManager {
	once.Do(func() {
		instance = &mongoManager{}
	})
	return instance
}

// GetMongoClient returns the MongoDB client instance
func (m *mongoManager) GetMongoAdapterClient() *mongo.Client {
	if m.client == nil {
		log.Fatal("MongoDB client is not initialized. Call InitializeMongoClient first.")
	}
	return m.client
}

// InitializeMongoClient initializes the MongoDB client
func (m *mongoManager) InitializeMongoAdapterClient() error {
	client, err := mongoAdapter.MongoService.GetClient()
	if err != nil {
		return err
	} else {
		m.client = client
	}
	return nil
}

// DisconnectMongoClient disconnects the MongoDB client
func (m *mongoManager) DisconnectMongoClient() error {
	if m.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := m.client.Disconnect(ctx)
		if err != nil {
			return err
		}
		m.client = nil
		log.Println("Disconnected from MongoDB!")
	}
	return nil
}

func GetDBAdapter() *adapters.MongoAdapter {
	port, _ := strconv.Atoi(os.Getenv("DB_HOST"))
	DSNMongoParam := entities.DSNMongo{
		Host:     os.Getenv("DB_URL"),
		Port:     port,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_COLLECTION"),
	}
	mongoAdapter = adapters.NewMongoAdapter(&DSNMongoParam)
	return mongoAdapter
}
