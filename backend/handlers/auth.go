package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "authentication-api/models"
    "authentication-api/utils"
)


func Login(c *gin.Context) {
	var user models.User

	if err := c.shouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return 
	}

	if user.Username == "user" && user.Password == "password" {
		token, err  := utils.GenerateToken(user.ID)
		

		if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return 
		}
        c.JSON(http.StatusOK, gin.H{"token": token})
	}

	else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func Register(c *gin.Context) {
	var user models.user

	if err := c.shouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return 
	}

	user.ID = 1
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}