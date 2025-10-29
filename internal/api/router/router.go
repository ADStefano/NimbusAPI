package router

import (
	"github.com/ADStefano/NimbusAPI/internal/api/router/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateRouter(log *zap.Logger) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	
	return routes.RegisterRoutes(log, router)

}
