package movie

import (
	"database/sql"
	"log"
	"sparrow-plus/types"

	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetMovies() ([]*types.Movie, error) {
	rows, err := s.db.Query(`SELECT * FROM "movies" WHERE "deletedAt" IS NULL`)

	if err != nil {
		return nil, err
	}

	movies := make([]*types.Movie, 0)
	for rows.Next() {
		movie, err := scanRowsIntoMovie(rows)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (s *Store) GetMovieById(movieId string) (*types.Movie, error) {
	rows, err := s.db.Query(`SELECT * FROM "movies" WHERE "deletedAt" IS NULL AND "movieId" = ?`, movieId)
	if err != nil {
		return nil, err
	}

	movie := new(types.Movie)
	for rows.Next() {
		movie, err = scanRowsIntoMovie(rows)
		if err != nil {
			return nil, err
		}
	}

	return movie, nil
}

func (s *Store) CreateMovie(movie types.CreateMoviePayload) error {
	_, err := s.db.Exec(
		"INSERT INTO movies (movieId, name, filePath) VALUES (?, ?, ?)",
		uuid.New(),
		movie.Name,
		movie.FilePath,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoMovie(rows *sql.Rows) (*types.Movie, error) {
	movie := new(types.Movie)
	err := rows.Scan(
		&movie.MovieId,
		&movie.Name,
		&movie.FilePath,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *Store) createMoviesTable() {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS "movies" (
			movieId    VARCHAR(36) PRIMARY KEY,
			name       VARCHAR(255) NOT NULL,
			filePath   VARCHAR(128) NOT NULL,
			categoryId VARCHAR(128) NOT NULL,
			createdAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deletedAt  TIMESTAMP
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store) SetupMovies() {
	s.createMoviesTable()
}
