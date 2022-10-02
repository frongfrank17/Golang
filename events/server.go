package main

import (
	db "events/database"
	routes "events/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Event Services")
	app := fiber.New()
	client := db.ConnectMongoDb("mongodb://localhost:27017")
	routes.PublicRoutes(app, client)
	app.Listen(":3000")

}
