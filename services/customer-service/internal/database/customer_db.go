package database

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"services/customer-service/internal/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/gift_shop_customer?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//Auto migrate tables
	err = db.AutoMigrate(&models.Customer{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Database connection established")
	
}

