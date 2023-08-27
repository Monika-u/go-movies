package config

import (
	"fmt"
	"time"

	mysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// PostgresDB : Variable to hold the postgresql db connection
var MariaDB *gorm.DB

// InitializeDB : Initializes the Database connections
func InitializeDB() {

	db, err := gorm.Open(mysql.Open(":8000"+fmt.Sprintf("?%s", "&parseTime=True")), &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            true,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true, TablePrefix: "movie" + "."},
		// Logger: Logger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	MariaDB = db
}
