package database

import (
	"book-comments/cmd/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:password@tcp(mysql:3306)/bookComments?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("err db conn : %v", err)
	}
	DB = db
}
func Migrate() {
	DB.AutoMigrate(&model.Category{}, &model.Author{}, &model.User{}, &model.Book{}, &model.Comment{}, &model.Ask{})
}
