package controllers

import (
	"EnvManager-api/database"
	"EnvManager-api/middleware"
	"EnvManager-api/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitializeUserRoutes(router *gin.Engine, client *database.DBClient) {
	dbClient = client
	router.POST("/admin/user", middleware.AdminOnly(), CreateUser)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := dbClient.InsertOne(ctx, "users", user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
