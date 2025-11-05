package router

import (
	"github.com/ADStefano/NimbusAPI/internal/api/router/routes"
	"github.com/ADStefano/NimbusAPI/internal/config"
	"github.com/gin-gonic/gin"
)

func CreateRouter(cfg *config.App) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	return routes.RegisterRoutes(cfg, router)

}
