package database

import (
	"log"
	"os"

	seller "gin/models/sellermodel"
	user "gin/models/usermodel"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DatabaseMigration() {
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loaing .env file")
	}
	urlDSN := os.Getenv("urlDSN")
	DB, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		panic("connection failed")
	}
	DB.AutoMigrate(&user.User{}, &seller.Collection{})
	Database = DB
}
