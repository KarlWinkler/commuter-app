package db

import (
	"database/sql"

	"github.com/karlwinkler/web-server/sqlc"
)

type QueryStore struct {
	*sqlc.Queries
	DB *sql.DB
}

func NewQueryStore(db *sql.DB) *QueryStore {
	return &QueryStore{
		DB: db,
		Queries: sqlc.New(db),
	}
}