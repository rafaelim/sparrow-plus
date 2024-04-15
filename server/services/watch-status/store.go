package watchstatus

import (
	"database/sql"
	"fmt"
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

func (s *Store) GetWatchStatus() ([]*types.WatchStatus, error) {
	rows, err := s.db.Query(`SELECT * FROM "watchStatus" WHERE "deletedAt" IS NULL`)

	if err != nil {
		return nil, err
	}

	watchStatuses := make([]*types.WatchStatus, 0)
	for rows.Next() {
		watchStatus, err := scanRowsIntoWatchStatus(rows)
		if err != nil {
			return nil, err
		}
		watchStatuses = append(watchStatuses, watchStatus)
	}

	return watchStatuses, nil
}

func (s *Store) UpsertWatchStatus(payload types.UpsertWatchStatusPayload) (*types.WatchStatus, error) {
	watchStatus := &types.WatchStatus{
		WatchStatusId: uuid.New().String(),
		RelationId:    payload.RelationId,
		RelationType:  payload.RelationType,
		Timestamp:     payload.Timestamp,
	}

	_, err := s.db.Exec(
		`
		INSERT INTO watchStatus (watchStatusId, relationId, relationType, timestamp)
			VALUES(?, ?, ?, ?)
		ON CONFLICT(relationId, relationType) DO UPDATE SET
		  relationId=excluded.relationId,
		  relationType=excluded.relationType,
		  timestamp=excluded.timestamp
		`,
		watchStatus.WatchStatusId,
		watchStatus.RelationId,
		watchStatus.RelationType,
		watchStatus.Timestamp,
	)
	fmt.Printf("%+v\n", watchStatus)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return watchStatus, nil
}

func scanRowsIntoWatchStatus(rows *sql.Rows) (*types.WatchStatus, error) {
	watchStatus := new(types.WatchStatus)
	err := rows.Scan(
		&watchStatus.WatchStatusId,
		&watchStatus.RelationId,
		&watchStatus.RelationType,
		&watchStatus.Timestamp,
		&watchStatus.CreatedAt,
		&watchStatus.UpdatedAt,
		&watchStatus.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return watchStatus, nil
}

func (s *Store) createWatchStatusTable() {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS "watchStatus" (
			watchStatusId VARCHAR(36) PRIMARY KEY,
			relationId    VARCHAR(36),
			relationType  VARCHAR(36),
			timestamp     INT NOT NULL,
			createdAt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deletedAt     TIMESTAMP,
			CONSTRAINT unique_relation UNIQUE (relationId, relationType)
		);
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store) SetupWatchStatus() {
	s.createWatchStatusTable()
}
