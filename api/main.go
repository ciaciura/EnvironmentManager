package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"EnvManager-api/config"
	"EnvManager-api/controllers"
	"EnvManager-api/database"
	docs "EnvManager-api/docs"
	"EnvManager-api/middleware"
)

var dbClient *database.DBClient

// @title EnvironmentManager
// @description Api for Environment access management.
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /api/v1

func main() {
	if config.Env == "DEV" {
		fmt.Println("Environment", config.Env)
		fmt.Println("JWT Secret:", config.JWT_SECRET)
		fmt.Println("MongoDB URI:", config.MONGODB_URI)
		fmt.Println("MongoDB Database:", config.MONGODB_DATABASE)
		fmt.Println("Admin Username:", config.ADMIN_USERNAME)
		fmt.Println("Admin Password:", config.ADMIN_PASSWORD)
		fmt.Println("App Port:", config.APP_PORT)
	}
	// Set up MongoDB client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}
	dbClient = &database.DBClient{Client: client, Database: config.MONGODB_DATABASE}

	// Ensure default admin user exists
	dbClient.EnsureDefaultAdmin(ctx)

	r := gin.Default()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Routes
	r.POST("/login", controllers.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/swagger", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Use(middleware.AuthMiddleware())
	controllers.InitializeServerRoutes(r, dbClient)
	controllers.InitializeUserRoutes(r, dbClient)
	// Run the server
	r.Run(config.APP_PORT)
}
