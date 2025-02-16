package service

import (
	"github.com/karlwinkler/web-server/sqlc"
)

type TripsResult struct {
	Count int
	Result []sqlc.Trip
}

type ModesResult struct {
	Count int
	Result []sqlc.Mode
}