package main

import (
	"github.com/gofiber/fiber/v2"
	ap "github.com/karlwinkler/web-server/api"
)

func main() {
	app := fiber.New()

	api := app.Group("/")
	api.Post("/", ap.Submit)

	app.Listen(":3001")
}