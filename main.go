package main

import (
	"github.com/Yadier01/rest-api/db"
	"github.com/Yadier01/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db.InitDB()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
