package controllers

import (
	"net/http"

	"GoCommerce/internal/db"
	"GoCommerce/internal/models"
	"GoCommerce/internal/repositories"
	"GoCommerce/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateCartItem(c *gin.Context) {
	cartRepo := repositories.NewCartRepository(db.DB)
	// Get the JSON body and decode into variables
	var cartItem models.CartItem
	if err := c.BindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return
	}

	// Create the cartItem
	if err := cartRepo.CreateCartItem(&cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not create the cartItem"})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"cartItem": cartItem})
}

func GetCartItems(c *gin.Context) {
	cartRepo := repositories.NewCartRepository(db.DB)
	cartItems, err := cartRepo.GetCartItems()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no cart items found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cartItems": cartItems})
}

func GetCartItemsByUserID(c *gin.Context) {
	cartRepo := repositories.NewCartRepository(db.DB)
	userID := c.Param("id")

	cartItems, err := cartRepo.GetCartItemsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no cart items found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"cartItems": cartItems})
}

func UpdateCartItem(c *gin.Context) {
	cartRepo := repositories.NewCartRepository(db.DB)
	var cartItem models.CartItem
	if err := c.BindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return
	}

	if err := cartRepo.UpdateCartItem(&cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not update the cart item"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"cartItem": cartItem})
}

func DeleteCartItem(c *gin.Context) {
	cartRepo := repositories.NewCartRepository(db.DB)
	itemID := utils.ParseUint(c.Param("id"))

	if err := cartRepo.DeleteCartItem(itemID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not delete the cart"})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func EmptyCart(c *gin.Context) {
	cartRepo := repositories.NewCartRepository(db.DB)
	userID := c.Param("id")

	if err := cartRepo.EmptyCart(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not empty the cart"})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
