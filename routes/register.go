package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Yadier01/rest-api/models"
	"github.com/gin-gonic/gin"
)

//this package is not for user singup or login, is for regsitering events

func registerForEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "that event does not exist"})
		return
	}
	var register models.Register

	register.EventID = event.ID
	register.UserID = c.MustGet("id").(int64)

	err = register.Save()
	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save registration"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})

}

func cancelRegistration(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse id"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "that event does not exist"})
		return
	}

	var register models.Register
	register.EventID = event.ID
	register.UserID = c.MustGet("id").(int64)

	err = register.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete registration"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "canceled"})
}
