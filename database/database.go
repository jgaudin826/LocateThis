package database

import (
	"log"

	"locate-this/database/dbmodel"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("LocateThis.db"),
		&gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	Migrate(DB)

	log.Println("Database connected and migrated")
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&dbmodel.User{},
		&dbmodel.Location{},
		&dbmodel.Group{},
		&dbmodel.GroupLocation{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")
}
