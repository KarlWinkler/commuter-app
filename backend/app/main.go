package main

import (
	"github.com/gofiber/fiber/v2"
	
	"github.com/karlwinkler/web-server/db"
	"github.com/karlwinkler/web-server/api"
	"github.com/karlwinkler/web-server/service"
	
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database := db.New()
	app := fiber.New()
	queryStore := db.NewQueryStore(database)
	handlers := service.NewHandler(queryStore)

	app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	router := app.Group("/")
	api.RegisterTripApi(router, handlers)
	api.RegisterModeApi(router, handlers)

	router.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	app.Listen(":3001")
}
