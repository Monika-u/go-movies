package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// // MariaDB : Variable to hold the MariaDB db connection
var MariaDB *gorm.DB

func InitializeDB() {
	// database connection info
	dsn := "username:password@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MariaDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}
