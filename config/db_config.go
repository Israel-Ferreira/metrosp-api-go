package config

import (
	"github.com/Israel-Ferreira/metrosp-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func SetupDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Db = db

	Db.AutoMigrate(&models.Line{}, &models.Station{})
}
