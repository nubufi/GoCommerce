package controllers

import (
	"net/http"

	"GoCommerce/internal/db"
	"GoCommerce/internal/models"
	"GoCommerce/internal/repositories"
	"GoCommerce/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	// Get the JSON body and decode into variables
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create the product
	if err := productRepo.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not create the product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func GetProducts(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	products, err := productRepo.GetProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Products not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProductByID(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	productID := utils.ParseUint(c.Param("id"))

	product, err := productRepo.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func UpdateProduct(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	productID := utils.ParseUint(c.Param("id"))

	// Get the JSON body and decode into variables
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update the product
	product.ID = productID
	if err := productRepo.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not update the product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func DeleteProduct(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	productID := utils.ParseUint(c.Param("id"))

	if err := productRepo.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not delete the product"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

