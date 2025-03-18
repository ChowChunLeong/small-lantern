package controller

import (
	"net/http"

	"github.com/ChowChunLeong/pineapple-language-api.git/database"
	"github.com/ChowChunLeong/pineapple-language-api.git/form"
	"github.com/ChowChunLeong/pineapple-language-api.git/repository"
	"github.com/gin-gonic/gin"
)

func OAuth(c *gin.Context){
	var request form.OAuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Call FindOrCreateUser function
	user, err := repository.FindOrCreateUser(database.Db[0], request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return user response
	c.JSON(http.StatusOK, gin.H{
		"message": "User authenticated successfully",
		"user":    user,
	})
}