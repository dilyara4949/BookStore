package models

import (
	"gorm.io/gorm"
	// "database/sql"
	_ "github.com/lib/pq"
	
)

type Book struct {
	gorm.Model
	Id int      			`json:"id"`
	Title string 			`json:"title"`
	Description string		`json:"description"`
	Cost float32			`json:"cost"`	
}

