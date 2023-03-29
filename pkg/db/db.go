package db

import (
	"log"

	"github.com/dilyara4949/BookStore/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:12345@localhost:5433/bookStore?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
			log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}