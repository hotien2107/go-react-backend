package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Println("Invalid id parameter")
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("Id is", id)

	movie, _ := app.models.DB.Get(id)

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()

	if err != nil {
		app.logger.Println("Invalid id parameter")
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")

	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.AllGenres()

	if err != nil {
		app.logger.Println("Invalid id parameter")
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")

	if err != nil {
		app.errorJSON(w, err)
		return
	}
}