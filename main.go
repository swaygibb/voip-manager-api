package main

import (
	"voip-manager-api/db"
	"voip-manager-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.InitDB()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.RegisterCallLogRoutes(r)
	routes.RegisterVoicemailRoutes(r)

	r.Run(":8080")
}
