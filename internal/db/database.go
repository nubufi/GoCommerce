package db

import (
	"gorm.io/gorm"
)

// Database is an interface that wraps basic database operations.
type Database interface {
	Create(value interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Where(query interface{}, args ...interface{}) Database
	Save(value interface{}) error
	Delete(value interface{}, conds ...interface{}) error
}

type gormDB struct {
	db *gorm.DB
}

// NewGormDB creates a new instance of gormDB that implements the Database interface.
func NewGormDB(db *gorm.DB) Database {
	return &gormDB{db}
}

// Create wraps GORM's Create method.
func (g *gormDB) Create(value interface{}) error {
	return g.db.Create(value).Error
}

// First wraps GORM's First method.
func (g *gormDB) First(dest interface{}, conds ...interface{}) error {
	return g.db.First(dest, conds...).Error
}

// Where wraps GORM's Where method.
func (g *gormDB) Where(query interface{}, args ...interface{}) Database {
	g.db = g.db.Where(query, args...)
	return g
}

// Save wraps GORM's Save method.
func (g *gormDB) Save(value interface{}) error {
	return g.db.Save(value).Error
}

// Delete wraps GORM's Delete method.
func (g *gormDB) Delete(value interface{}, conds ...interface{}) error {
	return g.db.Delete(value, conds...).Error
}
