package types

import "time"

type Category struct {
	CategoryId string     `json:"categoryId"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

type CategoryStore interface {
	GetCategories() ([]*Category, error)
}

type Movie struct {
	MovieId    string     `json:"movieId"`
	Name       string     `json:"name"`
	FilePath   string     `json:"filePath"`
	CategoryId string     `json:"categoryId"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

type CreateMoviePayload struct {
	MovieId    string `json:"movieId"`
	Name       string `json:"name"`
	FilePath   string `json:"filePath"`
	CategoryId string `json:"categoryId"`
}

type MovieStore interface {
	GetMovies() ([]*Movie, error)
	GetMovieById(movieId string) (*Movie, error)
	CreateMovie(movie CreateMoviePayload) error
}