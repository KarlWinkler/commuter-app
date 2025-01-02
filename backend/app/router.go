package main

import (
	"github.com/gofiber/fiber/v2"
	ap "github.com/karlwinkler/web-server/api"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:3030",
    AllowHeaders: "Origin, Content-Type, Accept",
}))

	api := app.Group("/")
	api.Post("/submit", ap.Submit)

	app.Listen(":3001")
}