package routes

import (
	"example.com/rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authGroup := server.Group("/")
	authGroup.Use(middleware.AuthMiddleware)
	{
		authGroup.POST("/events", createEvent)
		authGroup.PUT("/events/:id", updateEvent)
		authGroup.DELETE("/events/:id", deleteEvent)
	}

	server.POST("/register", registerUser)
	server.POST("/login", loginUser)
}
