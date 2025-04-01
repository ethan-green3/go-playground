package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"example.com/crypto-tracker/models" // update to match your go.mod name
)

var DB *gorm.DB

func InitDatabase() {

	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("DB HOST:", os.Getenv("DB_HOST"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.Coin{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Connected to database ✅")
}
