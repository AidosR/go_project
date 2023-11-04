package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&p.ID, &p.CreatedAt, &p.Version)
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
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
		WHERE id = $7 AND version = $8
		RETURNING version`

	args := []interface{}{
		p.Title,
		p.Description,
		p.Color,
		p.Material,
		p.Weight,
		p.Size,
		p.ID,
		p.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&p.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m PlayTentModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM play_tents
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
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

func (m PlayTentModel) GetAll(title string, color string, material string, filters Filters) ([]*PlayTent, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), *
		FROM play_tents
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (LOWER(color) = LOWER($2) OR $2 = '')
		AND (LOWER(material) = LOWER($3) OR $3 = '')
		ORDER BY %s %s, id ASC
		LIMIT $4 OFFSET $5`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{title, color, material, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	playTents := []*PlayTent{}

	for rows.Next() {
		var playTent PlayTent

		err := rows.Scan(
			&totalRecords,
			&playTent.ID,
			&playTent.CreatedAt,
			&playTent.Title,
			&playTent.Description,
			&playTent.Color,
			&playTent.Material,
			&playTent.Weight,
			&playTent.Size,
			&playTent.Version,
		)

		if err != nil {
			return nil, Metadata{}, err
		}

		playTents = append(playTents, &playTent)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return playTents, metadata, nil
}
