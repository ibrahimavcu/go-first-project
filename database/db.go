package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connection() *gorm.DB {
	if db == nil {
		dbUsername := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUsername, dbPassword, dbHost, dbPort, dbName)

		refDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Errorf("DB'ye bağlanılamadı.")
			return nil
		}

		sqlDB, err := refDB.DB()

		if err != nil {
			fmt.Errorf("sql.DB'e ulaşılamadı.")
			return nil
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(20)

		db = refDB
	}

	return db
}
