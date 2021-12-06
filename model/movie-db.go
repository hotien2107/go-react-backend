package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModels struct {
	DB *sql.DB
}

// Get a movie and error, if any
func (m *DBModels) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	// query := `select id, title, description, year, release_date, runtime, ratting, mpaa_rating, create_at, updated_at
	// 			from movies where id = $1`

	query := `select * from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreateAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `select mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movies_genres mg left join genres g
				on (g.id = mg.genre_id)
			where
				mg.movie_id = $1
	`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

	return &movie, nil
}

// Get all movies and error
func (m *DBModels) All() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `select * from movies order by title`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreateAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		query = `select mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movies_genres mg left join genres g
				on (g.id = mg.genre_id)
			where
				mg.movie_id = $1
`

		each_rows, err := m.DB.QueryContext(ctx, query, movie.ID)

		if err != nil {
			return nil, err
		}

		genres := make(map[int]string)
		for each_rows.Next() {
			var mg MovieGenre
			err := each_rows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			genres[mg.ID] = mg.Genre.GenreName
		}
		each_rows.Close()

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}


//get all genres
func (m *DBModels) AllGenres() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `select * from genres order by genre_name`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var genre Genre
		err := rows.Scan(
			&genre.ID,
			&genre.GenreName,
			&genre.CreateAt,
			&genre.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	return genres, nil
}