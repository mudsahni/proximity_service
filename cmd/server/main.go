package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"proximity_service_go/cmd/config"
	"proximity_service_go/internal/handler"
	"proximity_service_go/pkg/db"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration.", "Error", err)
	}

	// Connect to the MongoDB database
	err = db.Connect(cfg)
	if err != nil {
		slog.Error("Failed to connect to MongoDB.", "Error", err)
	}
	defer db.Disconnect()
	router := gin.Default()

	// Register routes
	handler.RegisterRoutes(router)

	// Start the server
	slog.Info("Server is running on port 8080")
	router.Run(":8080")

}
