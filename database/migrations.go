package database

import (
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Post{},
	)
	if err != nil {
		return err
	}
	return nil
}
