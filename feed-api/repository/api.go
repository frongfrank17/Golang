package repository

type ResponseAPI struct {
	Code int    `json:code`
	Data string `json:data`
}
type MongoAPI interface {
	GetProduct(id string) (*ResponseAPI, error)
	GetProducts() (*ResponseAPI, error)
}
