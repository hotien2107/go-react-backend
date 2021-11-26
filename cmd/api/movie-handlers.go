package main

import (
	models "backend/model"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Println("Invalid id parameter")
	}

	app.logger.Println("Id is", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Movie Title",
		Description: "Movie Description",
		Year:        2021,
		ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreateAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {

}
