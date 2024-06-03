package main

import (
	database "crowdfunding-backend/configs"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// loadEnv loads environment variables from a .env file.
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// connectDB establishes a connection to the database.
func connectDB() {
	if _, err := database.ConnectDB(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Connection to database is good")
}

func main() {
	loadEnv()
	connectDB()
}
