package main

import(
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func add(c *fiber.Ctx) error {
	a, errA := strconv.Atoi(c.Query("a", "0"))
	b, errB := strconv.Atoi(c.Query("b", "0"))
	if errA != nil || errB != nil {
		panic(fmt.Sprintf("%s%s", errA, errB))
	}
	sum := a + b
	return c.SendString(fmt.Sprintf("%d", sum))
}