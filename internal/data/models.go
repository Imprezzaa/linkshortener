package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Links LinkModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Links: LinkModel{DB: db},
	}
}
