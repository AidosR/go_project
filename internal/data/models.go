package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	PlayTents PlayTentModel
	Tokens    TokenModel
	Users     UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		PlayTents: PlayTentModel{DB: db},
		Tokens:    TokenModel{DB: db},
		Users:     UserModel{DB: db},
	}
}
