package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connect_db() {
	dsn := "host=localhost user=postgres dbname=go_blog port=5432 sslmode=disable timezone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failure")
	}

	database.AutoMigrate(&Post{})

	DB = database
}
