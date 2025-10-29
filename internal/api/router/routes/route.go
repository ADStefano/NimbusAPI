package routes

import (
	"github.com/ADStefano/NimbusAPI/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler gin.HandlerFunc
	Auth    bool
}

func RegisterRoutes(log *zap.Logger, router *gin.Engine) *gin.Engine {

	routes := CreateBucketRoutes()
	for _, route := range routes {
		if route.Auth {
			router.Handle(route.Method, route.Path, middleware.LoggerMiddleware(log), route.Handler)
		} else {
			router.Handle(route.Method, route.Path, middleware.LoggerMiddleware(log), route.Handler)
		}
	}
	return router
}
