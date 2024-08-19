package repositories

import (
	"GoCommerce/internal/db"
	"GoCommerce/internal/models"

	"gorm.io/gorm"
)

// CartRepository defines the interface for cart-related operations
type CartRepository interface {
	CreateCartItem(item *models.CartItem) error
	GetCartItems() ([]models.CartItem, error)
	GetCartItemsByUserID(userID string) ([]models.CartItem, error)
	UpdateCartItem(item *models.CartItem) error
	DeleteCartItem(itemID uint) error
	EmptyCart(userID string) error
}

type cartRepository struct {
	db *gorm.DB
}

// NewCartRepository creates a new cart repository
func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetCartItems() ([]models.CartItem, error) {
	cachedCartItems, err := db.GetCache("cartItems", []models.CartItem{})
	if err == nil {
		return cachedCartItems, nil
	}
	var cartItems []models.CartItem
	if err := r.db.Find(&cartItems).Error; err != nil {
		return nil, err
	}
	db.SetCache("cartItems", cartItems)
	return cartItems, nil
}

func (r *cartRepository) CreateCartItem(item *models.CartItem) error {
	db.ClearCache("cartItems")
	return r.db.Create(item).Error
}

func (r *cartRepository) UpdateCartItem(item *models.CartItem) error {
	db.ClearCache("cartItems")
	return r.db.Save(item).Error
}

func (r *cartRepository) GetCartItemsByUserID(userID string) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	if err := r.db.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return cartItems, err
	}
	return cartItems, nil
}

func (r *cartRepository) EmptyCart(userID string) error {
	db.ClearCache("cartItems")
	return r.db.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error
}

func (r *cartRepository) DeleteCartItem(itemID uint) error {
	db.ClearCache("cartItems")
	return r.db.Where("id = ?", itemID).Delete(&models.CartItem{}).Error
}
