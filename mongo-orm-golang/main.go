package main

import (
	"fmt"
	db "ormmongo/domain"
	"ormmongo/handle"
	"ormmongo/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("")
	client := db.ConnectMongoDb("mongodb://localhost:27017")
	repo := repository.NewProductRepository(client)
	controller := handle.NewProductHandler(repo)
	app.Get("/products/:id", controller.GetProduct)
	app.Get("/products", controller.GetProducts)
	app.Get("/products/film/size", controller.FilmSizeGroup)
	app.Get("/products/filmtype/:type", controller.FilmTypeGroup)
	app.Post("/products", controller.Created)
	app.Patch("/products/addprice/:id", controller.AddPrice)
	fmt.Println(client)
	app.Listen(":3000")
}
