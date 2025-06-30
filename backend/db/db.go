package db

import (
	"fmt"
	"log"
	"os"

	"github.com/dalhatmd/Missing-Child-Alert/user"
	"github.com/dalhatmd/Missing-Child-Alert/alert"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		
}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbuser, password, dbname)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = database.AutoMigrate(&user.User{}, &alert.Alert{})

	DB = database
	log.Println("Database connection established successfully\n Migrated successfully")

}

func GetDB() *gorm.DB {
	return DB
}
