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

	auth := app.Group("/auth")
	auth.POST("/signup", controllers.SignUp)
	auth.POST("/signin", controllers.SignIn)
	auth.GET("/signout", controllers.SignOut)
	
	order := app.Group("/order")
	order.POST("/create", controllers.CreateOrder)
	order.GET("/list", controllers.GetOrders)
	order.GET("/show/:id", controllers.GetOrderByID)
	order.GET("/show-by-user/:id", controllers.GetOrdersByUserID)
	order.PUT("/update/:id", controllers.UpdateOrder)
	order.DELETE("/delete/:id", controllers.DeleteOrder)

	app.Run()
}
