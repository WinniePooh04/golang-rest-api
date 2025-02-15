package utils

import "github.com/gin-gonic/gin"

func JSONRespone(c *gin.Context, e error, status int, message string, data interface{}) {
	if e != nil {
		c.JSON(500, gin.H{"message": "Unable to fetch " + message, "data": nil})
	} else {
		if data == nil {
			c.JSON(200, gin.H{"message": "Data Not Found", "data": data})
		} else {
			c.JSON(200, gin.H{"message": nil, "data": data})
		}
	}
}
