package controllers

import (
	"GoCommerce/internal/db"
	"GoCommerce/internal/models"
	"GoCommerce/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func SignUp(c *gin.Context) {
	// Get the JSON body and decode into variables
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Can not hash the password"})
	}

	// Create the user
	user.Password = string(hash)
	user.UserID = utils.GenerateRandomID()
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Can not create the user"})
	}

	setToken(c, user)

	c.JSON(200, gin.H{"user": user})
}

type body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
func SignIn(c *gin.Context) {
	// Get the JSON body and decode into variables
	var body body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}

	// Find the user
	var user models.User
	db.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
	}

	setToken(c, user)

	c.JSON(200, gin.H{"user": user})
}

func setToken(c *gin.Context, user models.User) {
	tokenString, err := utils.CreateToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to sign token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 60*60*24, "/", "", false, true)
}

func GetUserID(c *gin.Context) string {
	userID, _ := c.Get("userID")
	return userID.(string)
}

func SignOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/signin")
}

