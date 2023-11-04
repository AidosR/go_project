package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	PlayTents PlayTentModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		PlayTents: PlayTentModel{DB: db},
	}
}