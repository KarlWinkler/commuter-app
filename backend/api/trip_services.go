package api

import(
	"os"
	"fmt"
	"database/sql"

	"github.com/karlwinkler/web-server/sqlc"

	"github.com/gofiber/fiber/v2"
)

type CreateData struct {
	Name 		 string `json:"name"`
	Distance int    `json:"distance"`
	Mode 		 int    `json:"mode"`
}

func (h *Handlers) Submit(ctx *fiber.Ctx) error {
	fd := new(CreateData)
	if err := ctx.BodyParser(fd); err != nil {
		fmt.Fprintf(os.Stderr, "error processing body")
		return err
	}

	data := sqlc.CreateTripParams {
		Name: sql.NullString{String: fd.Name, Valid: true},
		Distance: int32(fd.Distance),
		Mode: sql.NullInt32{Int32: int32(fd.Mode), Valid: true},
	}

	trip, err := h.QueryStore.CreateTrip(ctx.Context(), data)
	if err != nil {
		return err
	}

	return ctx.JSON(fmt.Sprintf("Thank you for submitting, %s, %d, %d!", trip.Name, trip.Distance, trip.Mode))
}

func (h *Handlers) ListTrips(ctx *fiber.Ctx) error {
	trips, err := h.QueryStore.ListTrips(ctx.Context())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying database: %v\n", err)
		return err
	}

	result := TripsResult {
		Count: len(trips),
		Result: trips,
	}

	return ctx.JSON(result)
}
