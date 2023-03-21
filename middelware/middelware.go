package middelware

import (
	auth "gin/controller/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authrequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(auth.JwtKey), nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token expired"})
			panic(err)
		} else {
			// c.JSON(http.StatusAccepted, gin.H{"Message": "Token valid"})
			id := token.Claims.(jwt.MapClaims)["id"]
			role := token.Claims.(jwt.MapClaims)["role"]
			c.Set("id", id)
			c.Set("role", role)
			c.Next()
		}
	}
}
