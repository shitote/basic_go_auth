package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func COnnectDB() {
	dns := os.Getenv("DB_URL")
	var err error

	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
}
