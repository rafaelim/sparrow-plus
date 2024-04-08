package types

import "time"

type Episode struct {
	EpisodeId string     `json:"episodeId"`
	ShowId    string     `json:"showId"`
	Name      string     `json:"name"`
	Seasons   int        `json:"seasons"`
	FilePath  string     `json:"filePath"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type CreateEpisodePayload struct {
	EpisodeId string `json:"episodeId"`
	ShowId    string `json:"showId"`
	Name      string `json:"name"`
	Seasons   int    `json:"seasons"`
	FilePath  string `json:"filePath"`
}

type EpisodeStore interface {
	GetEpisodes(showId string, season int) ([]*Episode, error)
	GetEpisodeById(episodeId string) (*Episode, error)
	CreateEpisode(episode CreateEpisodePayload) error
}