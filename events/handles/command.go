package handles

import (
	"context"
	repositorys "events/repositorys"
	"net/http"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

type Commandhandles interface {
	Created(*fiber.Ctx) error
}
type commandhandles struct {
	repo repositorys.EventsRepo
}

func AdpterEventsCommandHandler(repo repositorys.EventsRepo) Commandhandles {
	return commandhandles{repo: repo}
}

func (event commandhandles) Created(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	type Request struct {
		Name       string `json:"name"`
		Start_time int64  `json:"start_time"`
	}
	req := new(Request)
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad!",
		})
	}
	err = event.repo.Created(ctx, req.Name, req.Start_time)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Created Successfuly !",
	})
}
