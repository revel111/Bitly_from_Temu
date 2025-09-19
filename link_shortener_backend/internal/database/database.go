package database

import (
	"fmt"
	"log"

	"linkShortener/configs"
	"linkShortener/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(data configs.ConfigData) (*gorm.DB, func(), error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		data.DbHost, data.DbUser, data.DbPass, data.DbName, data.DbPort)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	closer := func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting database: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}

	return db, closer, nil
}

func Migrate(db *gorm.DB) error {
	if db == nil {
		log.Fatalf("Database connection is not established")
	}

	return db.AutoMigrate(&model.Link{})
}
