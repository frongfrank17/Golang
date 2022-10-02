package routes

import (
	"events/handles"
	repository "events/repositorys"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func PublicRoutes(app *fiber.App, client *mongo.Client) {
	repo := repository.AdpterEventsRepository(client)
	controller := handles.AdpterEventsCommandHandler(repo)
	quericescontroller := handles.AdpterEventsQuericesHandler(repo)
	route := app.Group("/events")
	route.Post("/created", controller.Created)
	route.Get("/", quericescontroller.GetEvents)
	route.Get("/:id", quericescontroller.GetEventsOne)
}
