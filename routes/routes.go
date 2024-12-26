package routes

import (
	middleware "github.com/Yadier01/rest-api/middlwares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	auth := server.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.POST("/events", createEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)
	auth.POST("/events/:id/register", registerForEvent)
	auth.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/register", createUser)
	server.POST("/login", userLogin)
}
