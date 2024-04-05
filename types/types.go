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
