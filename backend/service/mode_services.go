package service

import(
	"os"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) ListModes(ctx *fiber.Ctx) error {
	modes, err := h.QueryStore.ListModes(ctx.Context())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying database: %v\n", err)
		return err
	}

	result := ModesResult {
		Count: len(modes),
		Result: modes,
	}

	return ctx.JSON(result)	
}
