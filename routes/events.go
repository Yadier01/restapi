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
		c.JSON(http.StatusNotFound, gin.H{"error": "could not get event"})
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
	event.UserID = c.MustGet("id").(int64)
	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created", "event": event})
}

func updateEvent(c *gin.Context) {

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

	if event.UserID != c.MustGet("id").(int64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not the owner of this event"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not parse data"})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not updated event "})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated", "event": updatedEvent})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "could not get event"})
	}
	// Check if the user is the owner of the event
	if event.UserID != c.MustGet("id").(int64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not the owner of this event"})
		return
	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not remove event"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "deleted"})
}
