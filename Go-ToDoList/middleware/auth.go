package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/NurochmanRizal/go-ToDoList/database"
	"github.com/NurochmanRizal/go-ToDoList/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *gin.Context) {
	tokenString, err:=c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(os.Getenv("SECRET")) , nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
    c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		database.DB.Find(&user, claims["sub"])
		if user.ID==0{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}	