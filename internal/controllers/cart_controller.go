package controllers

import (
	"net/http"

	"GoCommerce/internal/db"
	"GoCommerce/internal/models"
	"GoCommerce/internal/repositories"
	"GoCommerce/internal/utils"

	"github.com/gin-gonic/gin"
)

// CreateCartItem godoc
//	@Summary		Create a new cart item
//	@Description	Create a new cart item
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			cart_item	body		object{user_id=string,product_id=int,quantity=int,price=float64}	true	"Cart item details"
//	@Success		201			{object}	object{cartItem=object{id=uint,created_at=string,user_id=string,product_id=int,quantity=int,price=float64}}
//	@Failure		400			{object}	object{error=string}
//	@Failure		500			{object}	object{error=string}
//	@Router			/cart/create [post]
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

// GetCartItems godoc
//	@Summary		Get all cart items
//	@Description	Get all cart items
//	@Tags			Cart
//	@Produce		json
//	@Success		200	{object}	object{cartItems=[]object{id=uint,created_at=string,user_id=string,product_id=int,quantity=int,price=float64}}
//	@Failure		404	{object}	object{error=string}
//	@Router			/cart/list [get]
func GetCartItems(c *gin.Context) {
	cartRepo := repositories.NewCartRepository(db.DB)
	cartItems, err := cartRepo.GetCartItems()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no cart items found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cartItems": cartItems})
}

// GetCartItemsByUserID godoc
//	@Summary		Get all cart items by user ID
//	@Description	Get all cart items by user ID
//	@Tags			Cart
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	object{cartItems=[]object{id=uint,created_at=string,user_id=string,product_id=int,quantity=int,price=float64}}
//	@Failure		404	{object}	object{error=string}
//	@Router			/cart/show-by-user/{id} [get]
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

// UpdateCartItem godoc
//	@Summary		Update a cart item
//	@Description	Update a cart item
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int																	true	"Cart Item ID"
//	@Param			cart_item	body		object{user_id=string,product_id=int,quantity=int,price=float64}	true	"Cart item details"
//	@Success		200			{object}	object{cartItem=object{id=uint,created_at=string,user_id=string,product_id=int,quantity=int,price=float64}}
//	@Failure		400			{object}	object{error=string}
//	@Failure		500			{object}	object{error=string}
//	@Router			/cart/update/{id} [put]
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

// DeleteCartItem godoc
//	@Summary		Delete a cart item
//	@Description	Delete a cart item
//	@Tags			Cart
//	@Param			id	path	int	true	"Cart Item ID"
//	@Success		204
//	@Failure		500	{object}	object{error=string}
//	@Router			/cart/delete/{id} [delete]
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

// EmptyCart godoc
//	@Summary		Empty the cart
//	@Description	Empties the cart of a user
//	@Tags			Cart
//	@Param			id	path	string	true	"User ID"
//	@Success		204
//	@Failure		500	{object}	object{error=string}
//	@Router			/cart/empty/{id} [delete]
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
