package types

import "time"

type Episode struct {
	EpisodeId     string     `json:"episodeId"`
	ShowId        string     `json:"showId"`
	EpisodeNumber string     `json:"episodeNumber"`
	Name          string     `json:"name"`
	Season        string     `json:"season"`
	Path          string     `json:"path"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
}

type CreateEpisodePayload struct {
	EpisodeId     string `json:"episodeId"`
	ShowName      string `json:"showName"`
	ShowId        string `json:"-"`
	EpisodeNumber string `json:"episodeNumber"`
	Name          string `json:"name"`
	Season        string `json:"season"`
	Path          string `json:"path"`
}

type EpisodeStore interface {
	GetEpisodes(showId string) ([]*Episode, error)
	GetEpisodeById(episodeId string) (*Episode, error)
	CreateEpisode(episode CreateEpisodePayload) error
}
