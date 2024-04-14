package main

import (
	database "crowdfunding-backend/configs"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database using config
	_, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database is good")
}
