package main

import (
	"net/http"

	"github.com/Yadier01/rest-api/db"
	"github.com/Yadier01/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.Default()

	db.InitDB()
	srv.GET("/events", getEvents)
	srv.POST("/events", createEvent)
	srv.Run(":8080")
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get events"})
	}
	c.JSON(http.StatusOK, events)
}
func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not bind json"})
		return
	}
	event.UserID = 1
	event.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "created", "event": event})
}
