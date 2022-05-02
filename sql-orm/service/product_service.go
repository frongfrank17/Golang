package service

import (
	"fmt"
	"gorestdbstruc/repository"
)

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return productService{productRepo: productRepo}
}
func (s productService) Create(id string, name string, price int) (*ProductResponse, error) {

	Product := repository.Products{
		PRD_ID:    id,
		PRD_NAME:  name,
		PRD_PRICE: price,
	}
	err := s.productRepo.Insert(&Product)
	if err != nil {
		return nil, err
	}
	response := ProductResponse{
		PRODUCT_ID:   id,
		PRODUCT_NAME: name,
		PRICE:        price,
	}
	return &response, nil
}
func (s productService) GetProducts() ([]ProductResponse, error) {

	products, err := s.productRepo.GetAll()
	fmt.Println("get Data From Repository")
	fmt.Println(products)
	if err != nil {

		return nil, err
	}

	productResponses := []ProductResponse{}
	for _, product := range products {
		productResponse := ProductResponse{
			PRODUCT_ID:   product.PRD_ID,
			PRODUCT_NAME: product.PRD_NAME,
			PRICE:        product.PRD_PRICE,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}
func (s productService) GetProduct(id string) (*ProductResponse, error) {

	product, err := s.productRepo.GetONE(id)
	fmt.Println("get Data From Repository GET BY ID ")
	fmt.Println(product)
	if err != nil {

		return nil, err
	}

	proResponse := ProductResponse{
		PRODUCT_ID:   product.PRD_ID,
		PRODUCT_NAME: product.PRD_NAME,
		PRICE:        product.PRD_PRICE,
	}

	return &proResponse, nil
}
func (s productService) UpdatePrice(id string, price int) error {
	err := s.productRepo.UpdatePri(id, price)
	if err != nil {
		return err
	}
	return nil
}
