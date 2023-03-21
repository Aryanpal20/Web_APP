package auth

import (
	"fmt"
	"gin/database"
	user "gin/models/usermodel"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtToken struct {
	Token string `json:"token"`
}

var JwtKey = []byte("Jwt_Key")

func Register(c *gin.Context) {

	var input user.User
	var users user.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := user.User{Email: input.Email, Password: input.Password, Role: input.Role}

	if !strings.Contains(user.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Enter Email formate"})
		return
	}
	database.Database.Where("email = ?", user.Email).Find(&users)

	if users.Email != user.Email {

		password := []byte(string(user.Password))
		hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
		if err != nil {
			panic(err)
		}
		err = bcrypt.CompareHashAndPassword(hashedPassword, password)
		fmt.Println(err)
		user.Password = string(hashedPassword)
		database.Database.Create(&user)
		c.JSON(http.StatusAccepted, gin.H{"Message": "Registration succesfully !!!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email already exist"})
	}
}

func Login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	var users user.User

	database.Database.Where("email = ?", email).First(&users)

	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    users.ID,
			"email": users.Email,
			"role":  users.Role,
			"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		})
		tokenString, error := token.SignedString(JwtKey)
		c.JSON(http.StatusCreated, gin.H{"Token": tokenString})
		if error != nil {
			fmt.Println(error)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email and password"})
	}
}
