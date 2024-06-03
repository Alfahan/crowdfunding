package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := GetDSN()
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDSN() string {
	// Ambil nilai variabel lingkungan untuk pengguna, kata sandi, host, port, dan nama database
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	fmt.Printf("DB_USER: %s\n", user)
	fmt.Printf("DB_PASSWORD: %s\n", password)
	fmt.Printf("DB_HOST: %s\n", host)
	fmt.Printf("DB_PORT: %s\n", port)
	fmt.Printf("DB_NAME: %s\n", dbName)

	// Buat string DSN berdasarkan nilai variabel lingkungan yang diambil
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	return dsn
}
