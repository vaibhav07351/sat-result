package db

import (
	"log"
	"sat-result/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() *gorm.DB {

	dsn := "user=postgres password=5679 dbname=sat_result host=localhost port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto-migrate the schema (create table if not exists)
	err = Db.AutoMigrate(&model.Student{})
	if err != nil {
		log.Fatal("Failed to migrate the database schema:", err)
	}

	log.Println("Successfully connected to database")

	return Db
}
