package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Iheanacho-ai/authentication-api.git/initializers"
	"github.com/Iheanacho-ai/authentication-api.git/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func RequireAuth (c *gin.Context){
	//Collect the token from the cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//Check if token has expired

		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Check if there is an actual user attached to the token
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		} 

		// Send the user
		c.Set("user", user)
		//Continue from the middle ware
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	
	fmt.Println("Hey there whats up")
}
