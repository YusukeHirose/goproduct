package db

import (
	"fmt"
	"goproduct/api/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Manager() *gorm.DB {
	db := Connect()
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})
	return db
}

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(db:3306)/go_product?charset=utf8&parseTime=True&Local")
	if err != nil {
		log.Println("connection is faild")
		fmt.Println(err.Error())
	}
	db.LogMode(true)
	fmt.Println("db is connected")
	return db
}
