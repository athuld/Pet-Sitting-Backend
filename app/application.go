package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))
	mapUrls()
	router.Run(":8081")
}
