package service

import (
	"fmt"
	repository "redishex/repository"
)

type redisService struct {
	repo repository.ProductRepository
}

func NewRedisService(repo repository.ProductRepository) ProductServce {
	return redisService{repo: repo}
}
func (srv redisService) FeedData() ([]ProductReponse, error) {
	data, err := srv.repo.GetAll()
	if err != nil {
		return nil, err
	}
	products := []ProductReponse{}
	for _, prd := range data {
		products = append(products, ProductReponse{

			PRODUCT_NAME: prd.PRD_NAME,
			PRICE:        prd.PRD_PRICE,
		})
	}
	fmt.Println("service")
	return products, nil
}
func (srv redisService) CreateData(id string, name string, price int) ([]ProductReponse, error) {
	Product := repository.Products{
		PRD_ID:    id,
		PRD_NAME:  name,
		PRD_PRICE: price,
	}
	data, err := srv.repo.Insert(&Product)
	if err != nil {
		return nil, err
	}

	products := []ProductReponse{}
	for _, prd := range data {
		products = append(products, ProductReponse{

			PRODUCT_NAME: prd.PRD_NAME,
			PRICE:        prd.PRD_PRICE,
		})
	}
	fmt.Println("service Created")
	return products, nil
}
func (srv redisService) UpdateData() ([]ProductReponse, error) {

	return nil, nil
}
