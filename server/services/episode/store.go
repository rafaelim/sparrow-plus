package episode

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

func (s *Store) GetEpisodes(showId string) ([]*types.Episode, error) {
	rows, err := s.db.Query(`
		SELECT 
			*
		FROM "episodes" 
		WHERE 
			"showId" = ?
		`,
		showId,
	)

	if err != nil {
		return nil, err
	}

	episodes := make([]*types.Episode, 0)
	for rows.Next() {
		episode, err := scanRowsIntoEpisode(rows)
		if err != nil {
			return nil, err
		}
		episodes = append(episodes, episode)
	}

	return episodes, nil
}

func (s *Store) GetEpisodeById(episodeId string) (*types.Episode, error) {
	rows, err := s.db.Query(`SELECT * FROM "episodes" WHERE "deletedAt" IS NULL AND "episodeId" = ?`, episodeId)
	if err != nil {
		return nil, err
	}

	episode := new(types.Episode)
	for rows.Next() {
		episode, err = scanRowsIntoEpisode(rows)
		if err != nil {
			return nil, err
		}
	}

	return episode, nil
}

func (s *Store) CreateEpisode(episode types.CreateEpisodePayload) error {
	_, err := s.db.Exec(
		"INSERT INTO episodes (episodeId, showId, name, season, episodeNumber, path) VALUES (?, ?, ?, ?, ?, ?)",
		uuid.New(),
		episode.ShowId,
		episode.Name,
		episode.Season,
		episode.EpisodeNumber,
		episode.Path,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoEpisode(rows *sql.Rows) (*types.Episode, error) {
	episode := new(types.Episode)
	err := rows.Scan(
		&episode.EpisodeId,
		&episode.ShowId,
		&episode.Name,
		&episode.Season,
		&episode.EpisodeNumber,
		&episode.Path,
		&episode.CreatedAt,
		&episode.UpdatedAt,
		&episode.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return episode, nil
}

func (s *Store) createEpisodesTable() {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS "episodes" (
			episodeId     VARCHAR(36) PRIMARY KEY,
			showId 	      VARCHAR(36) NOT NULL,
			name          VARCHAR(255) NOT NULL,
			season        VARCHAR(36) NOT NULL,
			episodeNumber VARCHAR(36) NOT NULL,
			path          VARCHAR(128) NOT NULL,
			createdAt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deletedAt     TIMESTAMP,
			FOREIGN KEY(showId) REFERENCES shows(showId)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store) SetupEpisodes() {
	s.createEpisodesTable()
}
