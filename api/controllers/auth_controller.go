package controllers

import (
	"EnvManager-api/database"
	"EnvManager-api/middleware"
	"EnvManager-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var dbClient *database.DBClient

func Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Generate JWT Token
	tokenString, err := middleware.GenerateJWT(user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "role": user.Role})
}
