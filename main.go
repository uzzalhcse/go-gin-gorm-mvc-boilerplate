package main

import (
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/server"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	app.Boot()
	server.RunServer()

}
