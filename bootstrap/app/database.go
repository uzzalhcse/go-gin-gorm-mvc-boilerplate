package app

import (
	"fmt"
	"log"

	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/config"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

func createConnectionPool(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func buildConnectionString(dbConfig config.DBConfig) string {
	connString := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBDatabase,
		dbConfig.DBUsername,
		dbConfig.DBPassword,
	)
	return connString
}

func connectDB() error {
	rwConnString := buildConnectionString(config.Database.MasterDB)

	var err error
	DB, err = createConnectionPool(rwConnString)
	if err != nil {
		return fmt.Errorf("[INIT] failed to connect to the database: %v", err)
	}

	roConnString := buildConnectionString(config.Database.SlaveDB)

	err = DB.Use(
		dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				postgres.Open(roConnString),
			},
			Policy: dbresolver.RandomPolicy{},
		}),
	)
	if err != nil {
		return fmt.Errorf("[INIT] failed to configure database resolver: %v", err)
	}

	log.Println("[INIT] database connected")

	err = database.Migrate(DB)
	if err != nil {
		log.Printf("[INIT] failed/skipping migrating database: %v\n", err)
	}

	return nil
}
