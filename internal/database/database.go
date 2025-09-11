package database

import (
	"fmt"
	"linkShortener/configs"
	"linkShortener/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		configs.DbHost.GetValue(), configs.DbUser.GetValue(), configs.DbPass.GetValue(), configs.DbName.GetValue(), configs.DbPort.GetValue())

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return DB
}

func Migrate() {
	if DB == nil {
		log.Fatalf("Database connection is not established")
	}

	if err := DB.AutoMigrate(&model.Link{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
