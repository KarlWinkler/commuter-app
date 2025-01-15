package api

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterTripApi(router fiber.Router, handlers *Handlers) {
	router.Post("/submit", handlers.Submit)
	router.Get("/trips", handlers.List)
}
