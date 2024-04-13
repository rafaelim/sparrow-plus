package types

import "time"

type Movie struct {
	MovieId   string     `json:"movieId"`
	Name      string     `json:"name"`
	Year      string     `json:"year"`
	Path      string     `json:"path"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type CreateMoviePayload struct {
	MovieId string `json:"movieId"`
	Name    string `json:"name"`
	Year    string `json:"year"`
	Path    string `json:"path"`
}

type MovieStore interface {
	GetMovies() ([]*Movie, error)
	GetMovieById(movieId string) (*Movie, error)
	BulkCreateMovie(movies []CreateMoviePayload) error
	CreateMovie(movie CreateMoviePayload) error
}
