package main

import (
	"fmt"

	"github.com/Iheanacho-ai/authentication-api.git/controllers"
	"github.com/Iheanacho-ai/authentication-api.git/initializers"
	"github.com/Iheanacho-ai/authentication-api.git/middlewares"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run()


	fmt.Println("Hey there !")
}