package service

import (
	"github.com/karlwinkler/web-server/db"
)

type Handlers struct {
	QueryStore *db.QueryStore
}

func NewHandler(queryStore *db.QueryStore) *Handlers {
	return &Handlers{QueryStore: queryStore}
}