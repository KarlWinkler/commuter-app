package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	api := app.Group("/")
	api.Get("/", hello)
	api.Get("/greeting/:name", greeting)
	api.Get("/add", add)
	api.Post("/form", form)


	app.Listen(":3001")
}