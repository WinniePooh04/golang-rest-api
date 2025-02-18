package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	fttx := router.Group("fttx")
	{
		fttx.POST("/search", controllers.GetFTTx)
	}
}
