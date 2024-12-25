package routes

import (
	"net/http"
	"strconv"

	"github.com/Yadier01/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get events"})
		return
	}
	c.JSON(http.StatusOK, events)
}
func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not bind json"})
		return
	}
	event.UserID = 1

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created", "event": event})
}
