package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModels struct {
	DB *sql.DB
}

// Get a movie and error
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

	return &movie, nil
}

// Get all movies and error
func (m *DBModels) All(id int) ([]*Movie, error) {
	return nil, nil
}
