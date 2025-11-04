package main

import (
	"github.com/ADStefano/NimbusAPI/internal/api/router"
	"github.com/ADStefano/NimbusAPI/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	router := router.CreateRouter(cfg)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PETECO PETECO PETECO",
		})
	})

	router.Run(cfg.AppConfig.Port)

}
