package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func FTTxRoutes(r *gin.Engine) {
	r.POST("/search", controllers.GetFTTx)
}
