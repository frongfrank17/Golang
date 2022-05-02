package handle

import (
	"context"
	"fmt"
	"net/http"
	repository "ormmongo/repository"
	"reflect"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

type ProductHandle interface {
	GetProducts(*fiber.Ctx) error
	GetProduct(*fiber.Ctx) error
	Created(*fiber.Ctx) error
	FilmSizeGroup(*fiber.Ctx) error
	FilmTypeGroup(c *fiber.Ctx) error
	AddPrice(c *fiber.Ctx) error
}

type productHandle struct {
	repo repository.ProductRepo
}

func NewProductHandler(repo repository.ProductRepo) ProductHandle {
	return productHandle{repo: repo}
}

func (ph productHandle) GetProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Id := c.Params("id")
	defer cancel()
	fmt.Println("Handle , ", Id)
	result, err := ph.repo.FindOne(ctx, Id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(&result)

}
func (ph productHandle) GetProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := ph.repo.FindALL(ctx, 100, 0)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(&result)
}
func (ph productHandle) Created(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	type Request struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
		Type  string `json:"type"`
		Size  string `json:"size"`
	}
	req := new(Request)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad!",
		})
	}
	er := ph.repo.Created(ctx, req.Name, req.Price, req.Type, req.Size)
	if er != nil {
		return c.Status(http.StatusBadRequest).SendString(er.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Created Successfuly !",
	})
}
func (ph productHandle) FilmSizeGroup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := ph.repo.SizeFilm(ctx)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product Film !", "data": data[0],
	})
}
func (ph productHandle) FilmTypeGroup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	t := c.Params("type")
	data, err := ph.repo.TypeFilm(ctx, t)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product Film Type !", "data": data[0],
	})
}
func (ph productHandle) AddPrice(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Id := c.Params("id")
	PriceQ := c.Query("price")
	fmt.Println("Query Price ", PriceQ)
	fmt.Println(reflect.TypeOf(PriceQ))
	type Request struct {
		Price int `json:"price"`
	}
	Price := new(Request)

	if err := c.BodyParser(&Price); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad!",
		})
	}
	fmt.Println("HANDLE PRICE")
	fmt.Println(Price.Price)
	err := ph.repo.UpdatePrice(ctx, Id, Price.Price)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Add Price Product Successfuly !",
	})

}
