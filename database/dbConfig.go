package database

import (
	"github.com/gusrylmubarok/ism-api-golang/configuration"
	"github.com/gusrylmubarok/ism-api-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	/* Getting from dsn from .env there are
	** Host, User, Password, Database Name, Post Database Server
	*/
	dsn := configuration.DotEnv()
	
	// Make conntection into Driver Database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database.")
	}

	// Make Global Variabel to use ORM fiture
	DB = database
	
	// Auto Migration to database from schema model
	database.AutoMigrate(models.User{})
}