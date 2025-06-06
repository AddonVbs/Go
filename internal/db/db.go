package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Убираем AutoMigrate, т.к. миграции будут в ./migrations/*.sql
	// if err := db.AutoMigrate(&ts.Task{}); err != nil {
	//     log.Fatalf("Could not migrate: %v", err)
	// }

	return Db, nil
}
