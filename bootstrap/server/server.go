package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/providers"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/config"
	"log"
)

func RunServer() {
	if !config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	router := providers.RouteServiceProvider{}.Register(server)
	log.Fatal(router.Run(":" + config.App.Port))
}
