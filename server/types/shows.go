package types

import "time"

type Show struct {
	ShowId    string     `json:"showId"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type CreateShowPayload struct {
	ShowId string `json:"showId"`
	Name   string `json:"name"`
}

type ShowStore interface {
	GetShows() ([]*Show, error)
	GetShowById(showId string) (*Show, error)
	GetShowByName(name string) (*Show, error)
	CreateShow(show CreateShowPayload) (*Show, error)
}
