package controllers

import (
	"net/http"

	"GoCommerce/internal/db"
	"GoCommerce/internal/models"
	"GoCommerce/internal/repositories"
	"GoCommerce/internal/utils"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Create a new order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			order	body		object{user_id=string,total_price=float64,payment_method=string,shipping_address=string,order_items=[]object{product_id=int,quantity=float64,unit_price=float64,total_price=float64}}	true	"Order details"
//	@Success		201		{object}	object{id=uint,created_at=string,user_id=string,total_price=float64,payment_method=string,shipping_address=string,order_items=[]object{product_id=int,quantity=float64,unit_price=float64,total_price=float64}}
//	@Failure		400		{object}	object{error=string}
//	@Failure		500		{object}	object{error=string}
//	@Router			/order/create [post]
func CreateOrder(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	// Get the JSON body and decode into variables
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return
	}

	// Create the order
	if err := orderRepo.CreateOrder(&order, order.OrderItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not create the order"})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": order})
}

// GetOrders godoc
//
//	@Summary		Get all orders
//	@Description	Get all orders
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object{orders=[]object{id=uint,created_at=string,user_id=string,total_price=float64,payment_method=string,shipping_address=string,order_items=[]object{product_id=int,quantity=float64,unit_price=float64,total_price=float64}}}
//	@Failure		404	{object}	object{error=string}
//	@Router			/order/all [get]
func GetOrders(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	orders, err := orderRepo.GetOrders()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orders not found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetOrderByID godoc
//
//	@Summary		Get an order by ID
//	@Description	Get an order by ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		200	{object}	object{id=uint,created_at=string,user_id=string,total_price=float64,payment_method=string,shipping_address=string,order_items=[]object{product_id=int,quantity=float64,unit_price=float64,total_price=float64}}
//	@Failure		404	{object}	object{error=string}
//	@Router			/order/{id} [get]
func GetOrderByID(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	orderID := utils.ParseUint(c.Param("id"))

	order, err := orderRepo.GetOrderByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// GetOrdersByUserID godoc
//
//	@Summary		Get all orders by user ID
//	@Description	Get all orders by user ID
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	object{orders=[]object{id=uint,created_at=string,user_id=string,total_price=float64,payment_method=string,shipping_address=string,order_items=[]object{product_id=int,quantity=float64,unit_price=float64,total_price=float64}}}
//	@Failure		404	{object}	object{error=string}
//	@Router			/order/user/{id} [get]
func GetOrdersByUserID(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	userID := c.Param("id")

	orders, err := orderRepo.GetOrdersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orders not found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// UpdateOrder godoc
//
//	@Summary		Update an order
//	@Description	Update an order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int																																														true	"Order ID"
//	@Param			order	body		object{user_id=string,total_price=float64,payment_method=string,shipping_address=string,order_items=[]object{product_id=int,quantity=float64,unit_price=float64,total_price=float64}}	true	"Order details"
//	@Success		200		{object}	object{id=uint,created_at=string,user_id=string,total_price=float64,payment_method=string,shipping_address=string,order_items=[]object{product_id=int,quantity=float64,unit_price=float64,total_price=float64}}
//	@Failure		400		{object}	object{error=string}
//	@Failure		500		{object}	object{error=string}
//	@Router			/order/update [put]
func UpdateOrder(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	orderID := utils.ParseUint(c.Param("id"))
	
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return
	}
	
	order.ID = orderID
	if err := orderRepo.UpdateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not update the order"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// DeleteOrder godoc
//
//	@Summary		Delete an order
//	@Description	Delete an order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		204	{object}	object{}
//	@Failure		500	{object}	object{error=string}
//	@Router			/order/delete/{id} [delete]
func DeleteOrder(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	orderID := utils.ParseUint(c.Param("id"))

	if err := orderRepo.DeleteOrder(orderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not delete the order"})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
