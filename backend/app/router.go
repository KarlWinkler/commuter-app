package main

import (
	"github.com/gofiber/fiber/v2"
	ap "github.com/karlwinkler/web-server/api"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders: "Origin, Content-Type, Accept",
	}))
	// Initialize default config
	app.Use(logger.New())

	// Or extend your config for customization
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	api := app.Group("/")
	api.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api.Post("/submit", ap.Submit)

	app.Listen(":3001")
}