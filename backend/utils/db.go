package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("❌ DB_DSN is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	fmt.Println("✅ Connected to NeonDB!")

	DB = db
}
