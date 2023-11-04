package data

import (
	"go_project/internal/validator"
	"time"
)

type PlayTent struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Color       string    `json:"color,omitempty"`
	Material    string    `json:"material,omitempty"`
	Weight      Weight    `json:"weight,omitempty"`
	Size        string    `json:"size,omitempty"`
	Version     int32     `json:"version"`
}

func ValidatePlayTent(v *validator.Validator, tent *PlayTent) {
	v.Check(tent.Title != "", "title", "must be provided")
	v.Check(len(tent.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(len(tent.Description) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(len(tent.Color) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(len(tent.Material) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(len(tent.Size) <= 500, "title", "must not be more than 500 bytes long")
}
