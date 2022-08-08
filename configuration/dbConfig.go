package configuration

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connection() {
	dsn := DotEnv();

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database.")
	}

}