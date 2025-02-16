package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/karlwinkler/web-server/service"
)

func RegisterModeApi(router fiber.Router, handlers *service.Handlers) {
	router.Get("/modes", handlers.ListModes)
}
