package main

import (
	"go-rest-api/config"
	"go-rest-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load Env
	config.LoadEnv()

	// initialize DB
	services.InitDBAdapter()
	r.Run(":8080")
}
