package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/karlwinkler/web-server/service"
)

func RegisterTripApi(router fiber.Router, handlers *service.Handlers) {
	router.Post("/submit", handlers.Submit)
	router.Get("/trips", handlers.ListTrips)
}
