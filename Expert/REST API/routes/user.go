package routes

import (
	"net/http"

	"example.com/rest-api/models"

	"github.com/gin-gonic/gin"
)

func registerUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse user data" + err.Error()})
		return
	}

	err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	context.JSON(http.StatusCreated, user)
}
