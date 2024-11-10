package controllers

import (
	"context"
	"net/http"
	"time"

	"EnvManager-api/database"
	"EnvManager-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitializeServerRoutes(router *gin.Engine, client *database.DBClient) {
	dbClient = client
	router.GET("/servers", GetServers)
	router.POST("/servers/request", RequestAccess)
}

func GetServers(c *gin.Context) {
	collection := dbClient.GetCollection("servers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching servers"})
		return
	}
	defer cur.Close(ctx)

	var servers []models.Server
	for cur.Next(ctx) {
		var server models.Server
		if err := cur.Decode(&server); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding server data"})
			return
		}
		servers = append(servers, server)
	}

	c.JSON(http.StatusOK, servers)
}

func RequestAccess(c *gin.Context) {
	var serverRequest struct {
		ServerID primitive.ObjectID `json:"server_id"`
	}
	if err := c.BindJSON(&serverRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	accessRequest := models.AccessRequest{
		ServerID:  serverRequest.ServerID,
		Username:  username,
		Status:    "pending",
		Requested: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := dbClient.InsertOne(ctx, "access_requests", accessRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error requesting access"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access request submitted"})
}
