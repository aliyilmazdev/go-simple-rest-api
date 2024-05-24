package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aliyilmazdev/go-simple-rest-api/internal/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// Connect to database
	i, err := strconv.Atoi(os.Getenv("DB_PORT"))
    if err != nil {
        panic(err)
    }

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), i)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&product.Product{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}