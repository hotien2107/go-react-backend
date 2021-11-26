package models

import (
	"database/sql"
	"time"
)

//Models is wraper for database
type Models struct {
	DB DBModels
}

// NewModels return models with DB pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModels{DB: db},
	}
}

// Movie is the type for movie
type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"release_date"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreateAt    time.Time    `json:"create_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	MovieGenre  []MovieGenre `json:"-"`
}

// Genre is the type for genre
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Genre is the type for genre
type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	GenreID   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
