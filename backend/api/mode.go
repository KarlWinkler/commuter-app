package api

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterModeApi(router fiber.Router, handlers *Handlers) {
	router.Get("/modes", handlers.ListModes)
}
