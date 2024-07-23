package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "host=db user=user password=password dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Log{})
}

type Log struct {
	gorm.Model
	Request  string
	Response string
}

func logRequestResponse(request, response string) {
	log := Log{Request: request, Response: response}
	db.Create(&log)
}
