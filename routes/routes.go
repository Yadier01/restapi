package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", Middleware(), createEvent)
	server.PUT("/events/:id", Middleware(), updateEvent)
	server.DELETE("/events/:id", Middleware(), deleteEvent)

	server.POST("/register", createUser)
	server.POST("/login", userLogin)
}
