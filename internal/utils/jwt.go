package utils

import (
	"os"
	"time"

	"GoCommerce/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

// CreateToken creates a JWT token for a user
//
// Parameters:
//
// - user: The user for whom to create the token
//
// Returns:
//
// - string: The token string
//
// - error: An error if the token could not be created
func CreateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
