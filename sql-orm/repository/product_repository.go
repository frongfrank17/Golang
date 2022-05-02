package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) productRepository {
	return productRepository{db: db}
}
func (r productRepository) Insert(product *Products) error {

	err := r.db.Create(&product).Error
	if err != nil {

		return err
	}
	return nil
}
func (r productRepository) GetAll() ([]Products, error) {
	var products []Products
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r productRepository) GetONE(id string) (*Products, error) {
	var product Products
	fmt.Println(id)
	err := r.db.Where("prd_id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
func (r productRepository) UpdatePri(id string, price int) error {

	err := r.db.Model(Products{}).Where("prd_id = ?", id).Update("prd_price", price).Error

	if err != nil {
		return err
	}
	return nil
}
