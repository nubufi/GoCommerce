package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"GoCommerce/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectToDb connects to the postgresql database
func ConnectToDb() {
	var err error
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB")

	if host == "" {
		log.Fatal(errors.New("MYSQL_HOST environment variable is not set"))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai", user, password, host, port, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB.Debug()
	log.Println("Connected to database")
}

// Migrate migrates the models
func Migrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.CartItem{},
	)
}
