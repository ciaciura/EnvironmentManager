package controllers

import (
	"context"
	"net/http"
	"time"

	"EnvManager-api/database"
	"EnvManager-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func InitializeServerRoutes(router *gin.Engine, client *database.DBClient) {
	dbClient = client
	router.GET("/servers", GetServers)
	router.POST("/servers/request", RequestAccess)
	router.POST("/servers/add", AddServer)
}

// @Summary List all servers
// @Description Retrieve a list of all servers
// @Tags servers
// @Accept  json
// @Produce  json
// @Security JWT
// @Success 200 {array} models.Server
// @Failure 500 {object} map[string][]string "Error fetching servers"
// @Router /servers [get]
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

// @Summary Add a new server
// @Description Create a new server
// @Tags servers
// @Accept  json
// @Produce  json
// @Security JWT
// @Param server body models.Server true "Server details"
// @Success 200 {object} map[string][]string "Server created"
// @Failure 400 {object} map[string][]string "Invalid request"
// @Failure 500 {object} map[string][]string "Error creating a server"
// @Router /servers/add [post]
func AddServer(c *gin.Context) {
	var server models.Server
	if err := c.BindJSON(&server); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := dbClient.InsertOne(ctx, "servers", server)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating a server"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Server created"})
}

// @Summary Request access to a server
// @Description Submits an access request for a server
// @Tags servers
// @Accept  json
// @Produce  json
// @Security JWT
// @Param serverRequest body models.AccessRequestDTO true "Server access request details"
// @Success 200 {object} map[string][]string "Access request submitted"
// @Failure 400 {object} map[string][]string "Invalid request"
// @Failure 500 {object} map[string][]string "Error requesting access"
// @Router /servers/request-access [post]
func RequestAccess(c *gin.Context) {
	var serverRequest models.AccessRequestDTO
	if err := c.BindJSON(&serverRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	accessRequest := models.AccessRequest{
		ServerID:  serverRequest.ServerID,
		Username:  serverRequest.Username,
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
