package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"EnvManager-api/config"
	"EnvManager-api/controllers"
	"EnvManager-api/database"
	"EnvManager-api/middleware"
)

var dbClient *database.DBClient

func main() {
	// Set up MongoDB client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		log.Fatal(err)
	}
	dbClient = &database.DBClient{Client: client, Database: config.MongoDBDatabase}

	// Ensure default admin user exists
	dbClient.EnsureDefaultAdmin(ctx)

	r := gin.Default()

	// Routes
	r.POST("/login", controllers.Login)

	r.Use(middleware.AuthMiddleware())
	controllers.InitializeServerRoutes(r, dbClient)
	controllers.InitializeUserRoutes(r, dbClient)

	// Run the server
	r.Run(config.AppPort)
}
