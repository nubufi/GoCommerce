package main

import (
	"GoCommerce/internal/controllers"
	"GoCommerce/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	db.ConnectToDb()
	db.Migrate()
}

func main() {
	app := gin.New()

	app.POST("/signup", controllers.SignUp)
	app.POST("/signin", controllers.SignIn)
	app.GET("/signout", controllers.SignOut)

	app.Run()
}
