package db

import (
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitSQLServerDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_SQLSERVER_CONN_STR")
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
