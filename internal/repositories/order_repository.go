package repositories

import (
	"GoCommerce/internal/models"

	"gorm.io/gorm"
)

// OrderRepository defines the interface for order-related operations
type OrderRepository interface {
	CreateOrder(order *models.Order, items []models.OrderItem) error
	GetOrderByID(orderID uint) (*models.Order, error)
	GetOrdersByUserID(userID string) ([]models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(orderID uint) error
}

// orderRepository is the concrete implementation of the OrderRepository interface
type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// GetOrders retrieves all orders along with their associated order items
func (r *orderRepository) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// CreateOrder creates a new order and its associated order items
func (r *orderRepository) CreateOrder(order *models.Order, items []models.OrderItem) error {
	// Wrap the operation in a transaction to ensure atomicity
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// Insert the order items
		if err := tx.Create(&items).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetOrderByID retrieves an order by its ID along with its associated order items
func (r *orderRepository) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("OrderItems").First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// GetOrdersByUserID retrieves all orders for a specific user by their user ID
func (r *orderRepository) GetOrdersByUserID(userID string) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Where("user_id = ?", userID).Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrder updates an existing order in the database
func (r *orderRepository) UpdateOrder(order *models.Order) error {
	return r.db.Save(order).Error
}

// DeleteOrder deletes an order and its associated order items by its ID
func (r *orderRepository) DeleteOrder(orderID uint) error {
	// Wrap the operation in a transaction to ensure atomicity
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Delete the order items first
		if err := tx.Where("order_id = ?", orderID).Delete(&models.OrderItem{}).Error; err != nil {
			return err
		}

		// Delete the order itself
		if err := tx.Delete(&models.Order{}, orderID).Error; err != nil {
			return err
		}

		return nil
	})
}
