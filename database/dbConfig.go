package database

import (
	"github.com/gusrylmubarok/ism-api-golang/configuration"
	"github.com/gusrylmubarok/ism-api-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := configuration.DotEnv()

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database.")
	}

	DB = database

	database.AutoMigrate(models.User{})

}