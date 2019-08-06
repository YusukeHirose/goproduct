package db

import (
	"fmt"
	"log"

	"../api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DbManager() *gorm.DB {
	db := DbConnect()
	db.AutoMigrate(&models.Product{})

	return db
}

func DbConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(db:3306)/go_product?charset=utf8&parseTime=True&Local")
	if err != nil {
		log.Println("connection is faild")
		fmt.Println(err.Error())
	}
	db.LogMode(true)
	fmt.Println("db is connected")
	return db
}
