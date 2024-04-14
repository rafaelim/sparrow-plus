package show

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

func (s *Store) GetShows() ([]*types.Show, error) {
	rows, err := s.db.Query(`SELECT * FROM "shows" WHERE "deletedAt" IS NULL`)

	if err != nil {
		return nil, err
	}

	shows := make([]*types.Show, 0)
	for rows.Next() {
		show, err := scanRowsIntoShow(rows)
		if err != nil {
			return nil, err
		}
		shows = append(shows, show)
	}

	return shows, nil
}

func (s *Store) GetShowById(showId string) (*types.Show, error) {
	rows, err := s.db.Query(`SELECT * FROM "shows" WHERE "deletedAt" IS NULL AND "showId" = ?`, showId)
	if err != nil {
		return nil, err
	}

	show := new(types.Show)
	for rows.Next() {
		show, err = scanRowsIntoShow(rows)
		if err != nil {
			return nil, err
		}
	}

	return show, nil
}

func (s *Store) GetShowByName(name string) (*types.Show, error) {
	rows, err := s.db.Query(`SELECT * FROM "shows" WHERE "name" = ?`, name)
	if err != nil {
		return nil, err
	}

	show := new(types.Show)
	for rows.Next() {
		show, err = scanRowsIntoShow(rows)
		if err != nil {
			return nil, err
		}
	}

	return show, nil
}

func (s *Store) CreateShow(payload types.CreateShowPayload) (*types.Show, error) {
	show := &types.Show{
		ShowId: uuid.New().String(),
		Name:   payload.Name,
	}
	_, err := s.db.Exec(
		"INSERT INTO shows (showId, name) VALUES (?, ?)",
		show.ShowId,
		show.Name,
	)
	if err != nil {
		return nil, err
	}

	return show, nil
}

func scanRowsIntoShow(rows *sql.Rows) (*types.Show, error) {
	show := new(types.Show)
	err := rows.Scan(
		&show.ShowId,
		&show.Name,
		&show.CreatedAt,
		&show.UpdatedAt,
		&show.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return show, nil
}

func (s *Store) createShowsTable() {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS "shows" (
			showId     VARCHAR(36) PRIMARY KEY,
			name       VARCHAR(255) NOT NULL,
			createdAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deletedAt  TIMESTAMP
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store) SetupShows() {
	s.createShowsTable()
}
