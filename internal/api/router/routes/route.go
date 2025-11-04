package routes

import (
	"github.com/ADStefano/NimbusAPI/internal/api/middleware"
	"github.com/ADStefano/NimbusAPI/internal/config"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler gin.HandlerFunc
	Auth    bool
}

func RegisterRoutes(cfg *config.App, router *gin.Engine) *gin.Engine {

	routes := CreateBucketRoutes(*cfg.BucketConfig.Controller)
	for _, route := range routes {
		if route.Auth {
			router.Handle(route.Method, route.Path, middleware.LoggerMiddleware(cfg.Logger), route.Handler)
		} else {
			router.Handle(route.Method, route.Path, middleware.LoggerMiddleware(cfg.Logger), route.Handler)
		}
	}
	return router
}
