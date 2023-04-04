package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/controllers"
)

func InitRoutes() {
	PORT := os.Getenv("SERVER_PORT")
	gin.SetMode(os.Getenv("GIN_MODE"))

	health := new(controllers.HealthController)
	text := new(controllers.TextController)

	router := gin.Default()
	// logging to stdout
	router.Use(gin.Logger())
	log.Println("GIN_MODE = ", os.Getenv("GIN_MODE"))

	// recover from panics and respond with internal server error
	router.Use(gin.Recovery())

	// enabling cors
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// v1 group
	v1 := router.Group("/api/v1")

	// healthcheck
	healthGroup := v1.Group("/health")
	healthGroup.GET("", health.GetStatus)

	// text
	textGroup := v1.Group("/text")
	textGroup.GET("", text.GetTexts)
	textGroup.POST("", text.PostText)

	router.Run(":" + PORT)
}
