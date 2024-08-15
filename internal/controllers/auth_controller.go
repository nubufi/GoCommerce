package controllers

import (
	"net/http"

	"GoCommerce/internal/db"
	"GoCommerce/internal/models"
	"GoCommerce/internal/repositories"
	"GoCommerce/internal/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var userRepo = repositories.NewUserRepository(db.DB)
func SignUp(c *gin.Context) {
	// Get the JSON body and decode into variables
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not hash the password"})
		return
	}

	// Create the user
	user.Password = string(hash)
	user.UserID = utils.GenerateRandomID()
	

	if err := userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "This email is already taken"})
		return
	}

	setToken(c, user)

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

type body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context) {
	// Get the JSON body and decode into variables
	var body body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the user
	user, err := userRepo.FindUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	setToken(c, *user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func SignOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusNoContent, nil)
}

func DeleteAccount(c *gin.Context) {
	userID := GetUserID(c)

	if err := userRepo.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the user"})
		return
	}

	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusNoContent, nil)
}

func setToken(c *gin.Context, user models.User) {
	tokenString, err := utils.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 60*60*24, "/", "", false, true)
}

func GetUserID(c *gin.Context) string {
	userID, _ := c.Get("userID")
	return userID.(string)
}
