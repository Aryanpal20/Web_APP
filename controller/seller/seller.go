package seller

import (
	"fmt"
	"gin/database"
	seller "gin/models/sellermodel"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	ID        int       `json:"id"`
	Rupees    string    `json:"rupees"`
	Img       string    `json:"img_url"`
	Create_AT time.Time `json:"createdat"`
}

func SellerUploadImg(c *gin.Context) {

	var posts seller.Collection
	role := c.GetString("role")
	if role == "Seller" {
		rupees := c.PostForm("rupees")
		tokenid := c.GetFloat64("id")

		file, err := c.FormFile("file") // Get the uploaded file from the request
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		extension := filepath.Ext(file.Filename)
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Only JPG, JPEG, and PNG are allowed."})
			return
		}
		err = os.MkdirAll("uploads", os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		scheme1 := "http://"
		if c.Request.TLS != nil {
			scheme := "https://"
			posts.Image = scheme + c.Request.Host + "./uploads/" + file.Filename
		}
		posts.Image = "/uploads/" + file.Filename
		posts = seller.Collection{Image: posts.Image, Rupees: rupees, User_ID: int(tokenid), Createat: time.Now()}
		Image := scheme1 + c.Request.Host + "/static/" + file.Filename
		message := Message{ID: posts.User_ID, Img: Image, Rupees: rupees}

		err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		database.Database.Create(&posts)
		c.JSON(http.StatusCreated, gin.H{"Data": message})
		c.JSON(http.StatusOK, gin.H{"Message": "File uploaded successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not a seller"})
	}

}
