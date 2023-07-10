package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/middleware"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/routes"
)

type RouteServiceProvider struct{}

func (r RouteServiceProvider) Register(router *gin.Engine) *gin.Engine {
	router.Use(middleware.Cors)
	v1 := router.Group("/v1")
	{
		routes.SetupApiRoutes(v1)
	}

	routes.SetupWebRoutes(router)

	return router
}
