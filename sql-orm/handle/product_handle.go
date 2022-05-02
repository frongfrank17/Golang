package handle

import (
	"fmt"
	"gorestdbstruc/service"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	prdSrv service.ProductService
}

func NewProductHandler(prdSrv service.ProductService) productHandler {
	return productHandler{prdSrv: prdSrv}
}
func (h productHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.prdSrv.GetProducts()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(&products)
}
func (h productHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"message": " NOT ID PRODUCT"})
	}
	product, err := h.prdSrv.GetProduct(id)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(&product)
}
func (h productHandler) UpdatePrice(c *fiber.Ctx) error {

	type request struct {
		Id    string `json:"id"`
		Price int    `json:"price"`
	}

	req := new(request)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	if req.Id == "" {
		return c.Status(400).SendString("Bad Request")
	}
	err := h.prdSrv.UpdatePrice(req.Id, req.Price)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(fiber.Map{"Message": "Update Price"})
}
func (h productHandler) Create(c *fiber.Ctx) error {
	var product service.ProductResponse
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	chk, err_ := h.prdSrv.GetProduct(product.PRODUCT_ID)
	if err_ == nil {
		fmt.Println(chk)
		return c.Status(500).JSON(product)
	}
	newProduct, err := h.prdSrv.Create(product.PRODUCT_ID, product.PRODUCT_NAME, product.PRICE)
	if err != nil {
		return c.Status(500).SendString("Error Created Fail")
	}
	return c.Status(200).JSON(fiber.Map{"message": "Created Successfuly!", "data": newProduct})

}
