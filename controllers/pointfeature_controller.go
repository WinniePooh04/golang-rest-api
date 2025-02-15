package controllers

import (
	"go-rest-api/models"
	"go-rest-api/services"
	"go-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFTTx(c *gin.Context) {
	var fttxBody models.FTTxRequestBody
	if err := c.ShouldBindJSON(&fttxBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	points, err := services.GetPointFeatureswithFilter(fttxBody)
	utils.JSONRespone(c, err, http.StatusInternalServerError, "points", points)
}
