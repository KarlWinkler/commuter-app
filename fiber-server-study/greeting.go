package main

import(
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func greeting(c *fiber.Ctx) error {
	name := c.Params("name", "0")
	return c.SendString(fmt.Sprintf("Welcome, %s!", name))
}