package movie

import (
	"database/sql"
	"fmt"
	"log"
	"sparrow-plus/types"
	"strings"

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
	fmt.Println(movieId)
	rows, err := s.db.Query(`SELECT * FROM "movies" WHERE "movieId" = ?`, movieId)
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

func (s *Store) BulkCreateMovie(movies []types.CreateMoviePayload) error {
	if len(movies) == 0 {
		log.Println("No movie was inserted, received an empty array.")
		return nil
	}
	var insert = make([]string, 0, len(movies))
	var rowsArgs = make([]interface{}, 0, len(movies)*4)
	for _, movie := range movies {
		insert = append(insert, "(?, ?, ?, ?)")
		rowsArgs = append(rowsArgs, uuid.New())
		rowsArgs = append(rowsArgs, movie.Name)
		rowsArgs = append(rowsArgs, movie.Year)
		rowsArgs = append(rowsArgs, movie.Path)
	}
	query := fmt.Sprintf(`
		INSERT INTO "movies"
			(movieId, name, year, path)
		VALUES
			%v
	`, strings.Join(insert, ","))
	_, inerr := s.db.Exec(query, rowsArgs...)

	if inerr != nil {
		log.Fatal(inerr)
	}

	return nil
}

func (s *Store) CreateMovie(movie types.CreateMoviePayload) error {
	_, err := s.db.Exec(
		"INSERT INTO movies (movieId, name, year, path) VALUES (?, ?, ?, ?)",
		uuid.New(),
		movie.Name,
		movie.Year,
		movie.Path,
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
		&movie.Year,
		&movie.Path,
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
			year       VARCHAR(5) NOT NULL,
			path	   VARCHAR(128) NOT NULL,
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
