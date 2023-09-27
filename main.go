package main

import (
	"go_news/database"
	"go_news/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.DatabaseInit()
	database.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8000")

}
