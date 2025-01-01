package api

import(
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type FormData struct {
	Name string `json:"name"`
	Distance int `json:"distance"`
}

func Submit(c *fiber.Ctx) error {
	fd := new(FormData)
	if err := c.BodyParser(fd); err != nil {
		return err
	}
	
	return c.SendString(fmt.Sprintf("Thank you for submitting, %s!", fd.Name))
}