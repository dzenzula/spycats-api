package database

import (
	"cmd/catapi/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "host=localhost user=username password=password dbname=spycats port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to PostgreSQL database")

	err = DB.AutoMigrate(&models.SpyCat{}, &models.Mission{}, &models.Target{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	fmt.Println("Database migrations completed")
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Println("Error getting underlying SQL DB:", err)
			return
		}
		err = sqlDB.Close()
		if err != nil {
			log.Println("Error closing database connection:", err)
		}
	}
}
