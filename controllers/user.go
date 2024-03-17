package controllers

import (
	"net/http"

	"github.com/Iheanacho-ai/authentication-api.git/models"
	"github.com/Iheanacho-ai/authentication-api.git/initializers"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context){
	//collect data from body parameters

	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		}) 

		return

	}

	// Search if user exists already
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Email Address is alread in use",
		})

		return
	}
	
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	//create User
	user = models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)
	
	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, 
			gin.H{
				"error": "Failed to create user",
			})
	}
	
	//respond

	c.JSON(http.StatusOK, gin.H{
		"message": "User is created",
	})
}

func Login(){

}