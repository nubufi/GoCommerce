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

// SignUp godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		object{email=string,password=string,first_name=string,last_name=string}	true	"User details"
//	@Success		201		{object}	object{user=object{user_id=string,created_at=string,email=string,first_name=string,last_name=string,role=string}}
//	@Failure		400		{object}	object{error=string}
//	@Failure		409		{object}	object{error=string}
//	@Router			/auth/signup [post]
func SignUp(c *gin.Context) {
	userRepo := repositories.NewUserRepository(db.DB)
	// Get the JSON body and decode into variables
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not hash the password"})
		c.Abort()
		return
	}

	// Create the user
	user.Password = string(hash)
	user.UserID = utils.GenerateRandomID()

	if err := userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "This email is already taken"})
		c.Abort()
		return
	}

	setToken(c, user)

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

type body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignIn godoc
//
//	@Summary		Sign in
//	@Description	Signs in the user, returns the user details and sets the jwt token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		object{email=string,password=string}	true	"Login details"
//	@Success		200		{object}	object{user=object{user_id=string,email=string,created_at=string,first_name=string,last_name=string,role=string}}
//	@Failure		400		{object}	object{error=string}
//	@Failure		401		{object}	object{error=string}
//	@Failure		404		{object}	object{error=string}
//	@Router			/auth/signin [post]
func SignIn(c *gin.Context) {
	userRepo := repositories.NewUserRepository(db.DB)
	// Get the JSON body and decode into variables
	var body body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return
	}

	// Find the user
	user, err := userRepo.FindUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		c.Abort()
		return
	}

	setToken(c, *user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// SignOut godoc
//
//	@Summary		Sign out
//	@Description	Signs out the user
//	@Tags			Auth
//	@Produce		json
//	@Success		204
//	@Router			/auth/signout [get]
func SignOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusNoContent, nil)
}

// DeleteAccount godoc
//
//	@Summary		Delete the account
//	@Description	Deletes the account and signs out the user
//	@Tags			Auth
//	@Produce		json
//	@Success		204
//	@Failure		500	{object}	object{error=string}
//	@Router			/auth/delete [delete]
func DeleteAccount(c *gin.Context) {
	userRepo := repositories.NewUserRepository(db.DB)
	userID := GetUserID(c)

	if err := userRepo.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the user"})
		c.Abort()
		return
	}

	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusNoContent, nil)
}

func setToken(c *gin.Context, user models.User) {
	tokenString, err := utils.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
		c.Abort()
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 60*60*24, "/", "", false, true)
}

func GetUserID(c *gin.Context) string {
	userID, _ := c.Get("userID")
	return userID.(string)
}
