package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	Component string `json:"component" binding:"required"`
	Status    string `json:"status"`
}

func main() {
	router := gin.Default()

	router.POST("/healthCheck", func(c *gin.Context) {
		var h HealthCheck
		if err := c.ShouldBindJSON(&h); err != nil {
			h.Status = "error"
			c.JSON(400, h)

			return
		}
		log.Printf("I am OK: %v", h.Component)
		h.Status = "OK"
		c.JSON(200, h)
	})

	router.Run(":4000")
}
