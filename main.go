package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	router.Run(":4000")
}
