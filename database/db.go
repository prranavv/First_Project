package database

import (
	"log"

	"github.com/prranavv/first_project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBinstance struct {
	Db *gorm.DB
}

var Database DBinstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected into the database successfully")
	//Add Migrations
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Library{})
	Database = DBinstance{Db: db}
}
