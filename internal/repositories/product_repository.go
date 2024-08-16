package repositories

import (
	"GoCommerce/internal/models"

	"gorm.io/gorm"
)

// ProductRepository defines the interface for order-related operations
type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProducts() ([]models.Product, error)
	GetProductByID(productID uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID uint) error
}

// productRepository is the concrete implementation of the ProductRepository interface
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// GetProducts retrieves all products
func (r *productRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// CreateProduct creates a new order and its associated order items
func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

// GetProductByID retrieves an product by its ID
func (r *productRepository) GetProductByID(productID uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product
func (r *productRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

// DeleteProduct deletes an existing product
func (r *productRepository) DeleteProduct(productID uint) error {
	return r.db.Delete(&models.Product{}, productID).Error
}
