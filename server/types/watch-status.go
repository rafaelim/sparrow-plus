package types

import "time"

type WatchStatus struct {
	WatchStatusId string     `json:"watchStatusId"`
	RelationId    string     `json:"relationId"`
	RelationType  string     `json:"relationType"`
	Timestamp     int        `json:"timestamp"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
}

type UpsertWatchStatusPayload struct {
	WatchStatusId string `json:"watchStatusId"`
	RelationId    string `json:"relationId"`
	RelationType  string `json:"relationType"`
	Timestamp     int    `json:"timestamp"`
}

type WatchStatusStore interface {
	GetWatchStatus() ([]*WatchStatus, error)
	UpsertWatchStatus(watchStatus UpsertWatchStatusPayload) (*WatchStatus, error)
}
