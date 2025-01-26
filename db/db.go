package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// MySQL connection string: username:password@tcp(host:port)/dbname
	dsn := "root:1234567@tcp(127.0.0.1:3306)/go_mux?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established.")

	// Auto-migrate the models
	// err = DB.AutoMigrate(&models.Book{})
	// if err != nil {
	// 	log.Fatalf("Failed to migrate database: %v", err)
	// }
}
