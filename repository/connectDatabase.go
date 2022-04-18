package repository

import (
	"fmt"
	"github.com/kartikeya/sample_app/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=kartikeya dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(DB)
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&entity.Post{})
	fmt.Println("Database connected.......")
	return DB
}
