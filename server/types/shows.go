package types

import "time"

type Show struct {
	ShowId    string     `json:"showId"`
	Name      string     `json:"name"`
	Seasons   int        `json:"seasons"`
	FilePath  string     `json:"filePath"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type CreateShowPayload struct {
	ShowId   string `json:"showId"`
	Name     string `json:"name"`
	Seasons  int    `json:"seasons"`
	FilePath string `json:"filePath"`
}

type ShowStore interface {
	GetShows() ([]*Show, error)
	GetShowById(showId string) (*Show, error)
	CreateShow(show CreateShowPayload) error
}
