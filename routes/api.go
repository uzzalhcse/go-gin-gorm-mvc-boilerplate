package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/http/controllers"
)

func SetupApiRoutes(router *gin.RouterGroup) {
	var TestController = controllers.NewTestController()

	testRoutes := router.Group("/test")
	{
		testRoutes.GET("/", TestController.Test)
		testRoutes.GET("/bulk-insert", TestController.BulkInsert)
	}
}
