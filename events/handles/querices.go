package handles

import (
	"context"
	repositorys "events/repositorys"
	"net/http"
	"strconv"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

type Quericeshandles interface {
	GetEvents(*fiber.Ctx) error
	GetEventsOne(*fiber.Ctx) error
}
type quericeshandles struct {
	repo repositorys.EventsRepo
}

func AdpterEventsQuericesHandler(repo repositorys.EventsRepo) Quericeshandles {
	return quericeshandles{repo: repo}
}

func (event quericeshandles) GetEventsOne(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()
	result, err := event.repo.FindOne(ctx, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(&result)
}

func (event quericeshandles) GetEvents(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l := c.Query("limit")
	s := c.Query("skip")
	limit, er := strconv.ParseInt(l, 10, 64)
	if er != nil {
		return c.Status(http.StatusBadRequest).SendString(er.Error())
	}
	skip, ers := strconv.ParseInt(s, 10, 64)
	if ers != nil {
		return c.Status(http.StatusBadRequest).SendString(er.Error())
	}
	result, err := event.repo.FindAll(ctx, limit, skip)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(&result)
}
