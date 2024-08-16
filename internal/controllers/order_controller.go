package controllers

import (
	"net/http"

	"GoCommerce/internal/db"
	"GoCommerce/internal/models"
	"GoCommerce/internal/repositories"
	"GoCommerce/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	// Get the JSON body and decode into variables
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create the order
	if err := orderRepo.CreateOrder(&order, order.OrderItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not create the order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": order})
}

func GetOrders(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	orders, err := orderRepo.GetOrders()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orders not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetOrderByID(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	orderID := utils.ParseUint(c.Param("id"))

	order, err := orderRepo.GetOrderByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrdersByUserID(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	userID := c.Param("id")

	orders, err := orderRepo.GetOrdersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orders not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func UpdateOrder(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := orderRepo.UpdateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not update the order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func DeleteOrder(c *gin.Context) {
	orderRepo := repositories.NewOrderRepository(db.DB)
	orderID := utils.ParseUint(c.Param("id"))

	if err := orderRepo.DeleteOrder(orderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not delete the order"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
