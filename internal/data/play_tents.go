package data

import (
	"database/sql"
	"errors"
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

type PlayTentModel struct {
	DB *sql.DB
}

func (m PlayTentModel) Insert(p *PlayTent) error {
	query := `
		INSERT INTO play_tents (title, description, color, material, weight, size)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, version`

	args := []interface{}{p.Title, p.Description, p.Color, p.Material, p.Weight, p.Size}

	return m.DB.QueryRow(query, args...).Scan(&p.ID, &p.CreatedAt, &p.Version)
}

func (m PlayTentModel) Get(id int64) (*PlayTent, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT *
		FROM play_tents
		WHERE id = $1`

	var p PlayTent

	err := m.DB.QueryRow(query, id).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.Title,
		&p.Description,
		&p.Color,
		&p.Material,
		&p.Weight,
		&p.Size,
		&p.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &p, nil
}

func (m PlayTentModel) Update(p *PlayTent) error {
	query := `
		UPDATE play_tents
		SET title = $1, description = $2, color = $3, material = $4, weight = $5, size = $6, version = version + 1
		WHERE id = $7
		RETURNING version`

	args := []interface{}{
		p.Title,
		p.Description,
		p.Color,
		p.Material,
		p.Weight,
		p.Size,
		p.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&p.Version)
}

func (m PlayTentModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM play_tents
		WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
