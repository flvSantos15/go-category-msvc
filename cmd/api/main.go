package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // gin.Default() returns a pointer to a new gin router

	// define routes
	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	CategoryRoutes(router)

	// run the server
	router.Run(":8080")
}