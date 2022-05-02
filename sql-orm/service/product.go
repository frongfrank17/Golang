package service

type ProductResponse struct {
	PRODUCT_ID   string `json:"product_id"` // Explicitly specify the type to be uuid
	PRODUCT_NAME string `json:"product_name"`
	PRICE        int    `json:"product_price"`
}

//go:generate mockgen -destination=../mock/mock_service/mock_customer_service.go bank/service CustomerService
type ProductService interface {
	GetProducts() ([]ProductResponse, error)
	GetProduct(id string) (*ProductResponse, error)
	UpdatePrice(id string, price int) error
	Create(id string, name string, price int) (*ProductResponse, error)
}
