package models

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabasePostgreSQL() {
	dsn := "host=localhost user=gin password=gin dbname=gin port=5432 sslmode=disable TimeZone=Asia/Jakarta"
  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	database.AutoMigrate(&Product{})
	DB = database
}