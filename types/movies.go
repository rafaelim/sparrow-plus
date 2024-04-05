package types

import "time"

type Movie struct {
	MovieId   string     `json:"movieId"`
	Name      string     `json:"name"`
	FilePath  string     `json:"filePath"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type CreateMoviePayload struct {
	MovieId  string `json:"movieId"`
	Name     string `json:"name"`
	FilePath string `json:"filePath"`
}

type MovieStore interface {
	GetMovies() ([]*Movie, error)
	GetMovieById(movieId string) (*Movie, error)
	CreateMovie(movie CreateMoviePayload) error
}
