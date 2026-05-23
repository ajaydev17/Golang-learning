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

func loginUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse user data"})
		return
	}

	isValid, err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to validate credentials"})
		return
	}

	if !isValid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
