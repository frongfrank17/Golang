package handle

import (
	"encoding/json"
	repo "feddapi/repository"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

type ProductHandle interface {
	FindALL(*fiber.Ctx) error
	Find(*fiber.Ctx) error
}

func FindALL(c *fiber.Ctx) error {
	result, err := repo.GetProducts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	type GetProduct struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
		Ty    string `json:"type"`
		Si    string `json:"size"`
	}
	resultArr := []*GetProduct{}
	_ = json.Unmarshal([]byte(result.Data), &resultArr)
	type Resp struct {
		Product string `json:"product_name"`
		Price   int    `json:"price"`
	}
	respArr := []Resp{}
	for _, ra := range resultArr {
		Resp := Resp{
			Product: ra.Name,
			Price:   ra.Price,
		}
		respArr = append(respArr, Resp)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": respArr})
}

func Find(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := repo.GetProduct(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	type GetProduct struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
		Ty    string `json:"type"`
		Si    string `json:"size"`
	}
	resp := new(GetProduct)
	_ = json.Unmarshal([]byte(result.Data), &resp)
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": resp})

}
