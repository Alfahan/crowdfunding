package main

import (
	database "crowdfunding-backend/configs"
	"crowdfunding-backend/domains/users"
	"crowdfunding-backend/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	db, err := gorm.Open(mysql.Open(database.GetDSN()), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)

	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}
