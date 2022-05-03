package main

import (
	controller "feddapi/handle"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/api", controller.FindALL)
	app.Get("/api/:id", controller.Find)
	app.Listen(":3001")

}
