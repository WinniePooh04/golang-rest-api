package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {

	// Public Route (No authentication required)
	router.GET("/public", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a public endpoint"})
	})

	fttx := router.Group("fttx")
	fttx.Use(middlewares.PSKMiddleware()) // Apply pre-shared key middleware
	{
		fttx.POST("/search", controllers.GetFTTx)
	}
}
