package controllers

import (
	"net/http"

	"GoCommerce/internal/db"
	"GoCommerce/internal/models"
	"GoCommerce/internal/repositories"
	"GoCommerce/internal/utils"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	Create a new product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			product	body		object{name=string,description=string,price=float64}	true	"Product details"
//	@Success		201		{object}	object{id=uint,name=string,description=string,price=float64}
//	@Failure		400		{object}	object{error=string}
//	@Failure		500		{object}	object{error=string}
//	@Router			/product/create [post]
func CreateProduct(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	// Get the JSON body and decode into variables
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return
	}

	// Create the product
	if err := productRepo.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not create the product"})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

// GetProducts godoc
//
//	@Summary		Get all products
//	@Description	Get all products
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object{products=[]object{id=uint,name=string,description=string,price=float64}}
//	@Failure		404	{object}	object{error=string}
//	@Router			/product/all [get]
func GetProducts(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	products, err := productRepo.GetProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Products not found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProductByID godoc
//
//	@Summary		Get a product by ID
//	@Description	Get a product by ID
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	object{id=uint,name=string,description=string,price=float64}
//	@Failure		404	{object}	object{error=string}
//	@Router			/product/{id} [get]
func GetProductByID(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	productID := utils.ParseUint(c.Param("id"))

	product, err := productRepo.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Update a product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int														true	"Product ID"
//	@Param			product	body		object{name=string,description=string,price=float64}	true	"Product details"
//	@Success		200		{object}	object{id=uint,name=string,description=string,price=float64}
//	@Failure		400		{object}	object{error=string}
//	@Failure		500		{object}	object{error=string}
//	@Router			/product/{id} [put]
func UpdateProduct(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	productID := utils.ParseUint(c.Param("id"))

	// Get the JSON body and decode into variables
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		c.Abort()
		return
	}

	// Update the product
	product.ID = productID
	if err := productRepo.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not update the product"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Delete a product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		204	{object}	object{}
//	@Failure		500	{object}	object{error=string}
//	@Router			/product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	productRepo := repositories.NewProductRepository(db.DB)
	productID := utils.ParseUint(c.Param("id"))

	if err := productRepo.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not delete the product"})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
