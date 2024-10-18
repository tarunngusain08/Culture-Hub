package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/tarunngusain08/Culture-Hub/models"
)

func Connect() *gorm.DB {
	// Update with your PostgreSQL credentials
	dsn := "host=localhost dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Activity{})
	return db
}
