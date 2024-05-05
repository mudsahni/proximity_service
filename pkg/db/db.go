package db

import (
	"context"
	"fmt"
	"log/slog"
	"proximity_service_go/cmd/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect(cfg *config.Config) error {
	// Create a new MongoDB client
	mongoURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	if cfg.DBUser != "" && cfg.DBPassword != "" {
		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort)
	}
	clientOptions := options.Client().ApplyURI(mongoURI).SetMaxPoolSize(100).SetMinPoolSize(10)

	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to create MongoDB: %s", err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error("failed to ping MongoDB.", "error", err)
		return fmt.Errorf("failed to ping MongoDB: %s", err)
	}

	slog.Info("Connected to MongoDB!")
	return nil
}

func GetClient() *mongo.Client {
	return client
}

func Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		slog.Error("failed to disconnect from MongoDB.", "Error", err)
		return fmt.Errorf("failed to disconnect from MongoDB: %s", err)
	}
	slog.Info("Disconnected from MongoDB!")
	return nil
}
