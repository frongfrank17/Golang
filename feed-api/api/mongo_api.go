package api

import (
	"io/ioutil"
	"net/http"
)

const URL = "http://localhost:3000/"

type MongoAPI interface {
	GetProduct(id string) (*ResponseAPI, error)
	GetProducts() (*ResponseAPI, error)
}

type ResponseAPI struct {
	code int    `json:code`
	data string `json:data`
}

const Client = &http.Client{}

func GetProduct(id string) (*ResponseAPI, error) {
	uri := URL + "products/" + id
	request, err := http.Get(uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http.Header{

		"Content-Type": []string{"application/json"},
	}
	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	Code := response.StatusCode
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result := ResponseAPI{
		code: Code,
		data: string(data),
	}
	return result, nil
}
func GetProducts() (*ResponseAPI, error) {
	uri := URL + "products/"
	request, err := http.Get(uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http.Header{

		"Content-Type": []string{"application/json"},
	}
	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	Code := response.StatusCode
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result := ResponseAPI{
		code: Code,
		data: string(data),
	}
	return result, nil
}
