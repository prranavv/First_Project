package database

import (
	"log"

	"github.com/prranavv/first_project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	Db *gorm.DB
}

var Database DBinstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected into the database successfully")
	//Add Migrations
	db.AutoMigrate(&models.Book{}, &models.Library{})
	Database = DBinstance{Db: db}
}
