package config

import (
	"log"

	"github.com/noskine/pilot_api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db *gorm.DB
}

func NewConnection() *Database {
	dbURL := NewEnvironments().GetDatabaseUri()

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Failed to connect database")
	}

	if err := db.AutoMigrate(&models.Publications{}); err != nil {
		log.Fatal("Failed loading migrations")
	}

	return &Database{
		Db: db,
	}
}
