package main

import (
	"fmt"
	"go_project/internal/data"
	"go_project/internal/validator"
	"net/http"
	"time"
)

func (app *application) createPlayTentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string      `json:"title"`
		Description string      `json:"description"`
		Color       string      `json:"color"`
		Material    string      `json:"material"`
		Weight      data.Weight `json:"weight"`
		Size        string      `json:"size"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	tent := &data.PlayTent{
		Title:       input.Title,
		Description: input.Description,
		Color:       input.Color,
		Material:    input.Material,
		Weight:      input.Weight,
		Size:        input.Size,
	}

	v := validator.New()

	if data.ValidatePlayTent(v, tent); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showPlayTentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	playTent := data.PlayTent{
		ID:          id,
		CreatedAt:   time.Now(),
		Title:       "Play Tent",
		Description: "good play tent",
		Color:       "red",
		Material:    "polyester",
		Weight:      1.5,
		Size:        "for 2-3 children",
		Version:     1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"play_tent": playTent}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
