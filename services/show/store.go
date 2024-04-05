package show

import (
	"database/sql"
	"log"
	"sparrow-plus/types"
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

func (s *Store) CreateShow(show types.CreateShowPayload) error {
	_, err := s.db.Exec(
		"INSERT INTO shows (name seasons, filePath, categoryId) VALUES (?, ?, ?, ?)",
		show.Name,
		show.Seasons,
		show.FilePath,
		show.CategoryId,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoShow(rows *sql.Rows) (*types.Show, error) {
	show := new(types.Show)
	err := rows.Scan(
		&show.ShowId,
		&show.Name,
		&show.Seasons,
		&show.FilePath,
		&show.CategoryId,
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
		seasons    INT NOT NULL,
		filePath   VARCHAR(128) NOT NULL,
		categoryId VARCHAR(128) NOT NULL,
		createdAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deletedAt  TIMESTAMP,
		FOREIGN KEY(categoryId) REFERENCES categories(categoryId)
	)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store) SetupShows() {
	s.createShowsTable()
}
