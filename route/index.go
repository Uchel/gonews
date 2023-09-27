package route

import (
	"go_news/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handler.UserHandlerGetAll)
	r.Post("/", handler.UserHandlerCreate)
	r.Get("/:id", handler.UserHandlerGetById)
	r.Put("/:id", handler.UserHandlerUpdateById)
}
