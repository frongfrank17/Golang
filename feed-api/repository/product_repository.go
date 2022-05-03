package repository

import (
	"io/ioutil"
	"net/http"
)

const URL = "http://localhost:3000/"

func GetProduct(id string) (*ResponseAPI, error) {
	client := &http.Client{}
	uri := URL + "products/" + id
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http.Header{

		"Content-Type": []string{"application/json"},
	}
	response, err := client.Do(request)
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
		Code: Code,
		Data: string(data),
	}
	return &result, nil
}
func GetProducts() (*ResponseAPI, error) {
	client := &http.Client{}
	uri := URL + "products"

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http.Header{

		"Content-Type": []string{"application/json"},
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	Code := response.StatusCode
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	Tosrt := string(data)
	result := ResponseAPI{
		Code: Code,
		Data: Tosrt,
	}
	return &result, nil
}
