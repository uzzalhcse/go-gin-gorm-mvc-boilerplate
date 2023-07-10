package app

import (
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/config"
	"log"
)

func Boot() {
	config.Boot()
	err := connectDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
