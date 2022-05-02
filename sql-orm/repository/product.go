package repository

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model        // Adds some metadata fields to the table
	PRD_ID     string `json:"product_id"` // Explicitly specify the type to be uuid
	PRD_NAME   string `json:"product_name"`
	PRD_PRICE  int    `json:"product_price"`
}

type ProductRepository interface {
	GetAll() ([]Products, error)
	GetONE(id string) (*Products, error)
	UpdatePri(id string, price int) error
	Insert(product *Products) error
}
