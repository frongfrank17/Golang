package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type redisRepositoryRedis struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRedisRepositoryRedis(db *gorm.DB, redis *redis.Client) ProductRepository {
	return redisRepositoryRedis{db, redis}
}

func (r redisRepositoryRedis) GetAll() ([]Products, error) {
	var products []Products
	key := "repository::GetProducts"
	productsJson, err := r.redis.Get(context.Background(), key).Result()

	// GET DATA FORM REDIS
	if err == nil {

		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	err = r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	err = r.redis.Set(context.Background(), key, string(data), 0).Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r redisRepositoryRedis) GetONE(id string) (*Products, error) {
	var product Products
	fmt.Println(id)
	err := r.db.Where("prd_id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r redisRepositoryRedis) Insert(product *Products) ([]Products, error) {
	key := "repository::GetProducts"
	var products []Products
	err := r.db.Create(&product).Error
	fmt.Println("Insert")
	if err != nil {
		fmt.Println("Error Create")
		return nil, err
	}

	err = r.db.Find(&products).Error
	if err != nil {

		return nil, err
	}
	data, err := json.Marshal(products)

	err = r.redis.Set(context.Background(), key, string(data), 0).Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}
