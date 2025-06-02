package db

import (
	ts "BackEnd/internal/taskservice"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	if err := /*db.AutoMigrate*/ (&ts.Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	return db, nil
}
