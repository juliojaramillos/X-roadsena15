package app

import (
	"github.com/gin-gonic/gin"

	"unal.edu.co/rest-example/api/routing"
)

func mapRoutes() {
	// Ping route, health control.
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// API Version 1 group (al routes should start with /v1/*.
	v1 := router.Group("")

	// Map entities routing.
	routing.UserRouting(v1)
}
