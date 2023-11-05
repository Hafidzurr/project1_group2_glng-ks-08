package database

import (
	"github.com/Hafidzurr/project1_group2_glng-ks-08/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init(connectionString string) *gorm.DB {
	var err error
	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB.AutoMigrate(&models.Todo{})
	return DB
}
