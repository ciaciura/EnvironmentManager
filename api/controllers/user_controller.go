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

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Security JWT
// @Param user body models.User true "User details"
// @Success 200 {object} map[string][]string "User created"
// @Failure 400 {object} map[string][]string "Invalid request"
// @Failure 500 {object} map[string][]string "Error creating user"
// @Router /admin/user [post]
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
