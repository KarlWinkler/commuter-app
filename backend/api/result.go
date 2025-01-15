package api

import (
	"github.com/karlwinkler/web-server/sqlc"
)

type TripsResult struct {
	Count int
	Result []sqlc.Trip
}