package main

import (
	"errors"
	"fmt"
	"go_project/internal/data"
	"go_project/internal/validator"
	"net/http"
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

	playTent := &data.PlayTent{
		Title:       input.Title,
		Description: input.Description,
		Color:       input.Color,
		Material:    input.Material,
		Weight:      input.Weight,
		Size:        input.Size,
	}

	v := validator.New()

	if data.ValidatePlayTent(v, playTent); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.PlayTents.Insert(playTent)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/play_tents/%d", playTent.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"play_tent": playTent}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showPlayTentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	playTent, err := app.models.PlayTents.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"play_tent": playTent}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePlayTentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	playTent, err := app.models.PlayTents.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title       *string      `json:"title"`
		Description *string      `json:"description"`
		Color       *string      `json:"color"`
		Material    *string      `json:"material"`
		Weight      *data.Weight `json:"weight"`
		Size        *string      `json:"size"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		playTent.Title = *input.Title
	}
	if input.Description != nil {
		playTent.Description = *input.Description
	}
	if input.Color != nil {
		playTent.Color = *input.Color
	}
	if input.Material != nil {
		playTent.Material = *input.Material
	}
	if input.Weight != nil {
		playTent.Weight = *input.Weight
	}
	if input.Size != nil {
		playTent.Size = *input.Size
	}

	v := validator.New()
	if data.ValidatePlayTent(v, playTent); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.PlayTents.Update(playTent)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"plat_tent": playTent}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deletePlayTentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.PlayTents.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "play tent successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
