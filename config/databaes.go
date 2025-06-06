package config

import (
	"E_commerce_System/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "Manager:password@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect database: " + err.Error())
	}
	if err := db.AutoMigrate(&models.Product{}); err != nil {
		panic("auto shift fail: " + err.Error())
	}
	ERROR := db.AutoMigrate(&models.User{})
	if ERROR != nil {
		panic("Create table fail: " + ERROR.Error())
	}
	DB = db
	fmt.Println("Connected database successfully")
}
