package controllers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type TestController struct {
	db *gorm.DB
}

func NewTestController() *TestController {
	return &TestController{
		db: app.DB,
	}
}

func (ctrl TestController) Test(c *gin.Context) {
	var items []models.Post
	ctrl.db.Order("id DESC").Limit(500).Find(&items)
	c.JSON(200, gin.H{
		"message": "Test Success",
		"data":    items,
	})
}

func (ctrl TestController) BulkInsert(c *gin.Context) {
	for i := 0; i < 2000; i++ {
		test := models.Post{
			Name:        faker.Name(),
			Slug:        faker.Username(),
			Description: faker.Paragraph(),
		}

		ctrl.db.Create(&test)
	}
	var items []models.Post

	// Query the last 100 records from the database
	ctrl.db.Order("id DESC").Limit(500).Find(&items)

	c.JSON(200, gin.H{
		"message": "Test Success",
		"data":    items,
	})
}
