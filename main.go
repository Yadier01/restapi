package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.Default()

	srv.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	srv.Run(":8080")
}
