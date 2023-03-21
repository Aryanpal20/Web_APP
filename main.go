package main

import (
	auth "gin/controller/auth"
	"gin/controller/seller"
	"gin/database"
	"gin/middelware"

	"github.com/gin-gonic/gin"
)

func main() {

	database.DatabaseMigration()

	r := gin.Default()

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.POST("/sellerpost", middelware.Authrequired(), seller.SellerUploadImg)

	r.Run()
}
