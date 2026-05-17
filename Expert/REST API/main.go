package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event data" + err.Error()})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
