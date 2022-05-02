package main

import (
	"gorestdbstruc/configs"
	"gorestdbstruc/handle"
	"gorestdbstruc/repository"
	"gorestdbstruc/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
)

const DNS = "root:root@tcp(localhost:3333)/DBProduct?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	app := fiber.New()
	dial := mysql.Open(DNS)
	db, err := configs.DatabaseInit(dial)
	if err != nil {
		panic("Database Error")
	}
	proRepo := repository.NewProductRepositoryDB(db)
	proSer := service.NewProductService(proRepo)
	proHan := handle.NewProductHandler(proSer)
	app.Get("/product", proHan.GetProducts)
	app.Get("/product/:id", proHan.GetProduct)
	app.Post("/product/update", proHan.UpdatePrice)
	app.Post("/product", proHan.Create)
	app.Listen(":3000")
}
