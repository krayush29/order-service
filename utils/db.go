package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"order-service/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "user=krayush29 password=password12 dbname=orderdb sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate the Order model
	DB.AutoMigrate(&models.Order{})
}
