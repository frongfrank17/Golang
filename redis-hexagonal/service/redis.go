package service

type ProductReponse struct {
	PRODUCT_NAME string `json:"product_name"`
	PRICE        int    `json:"product_price"`
}

type ProductServce interface {
	FeedData() ([]ProductReponse, error)
	CreateData(id string, name string, price int) ([]ProductReponse, error)
	UpdateData() ([]ProductReponse, error)
}
