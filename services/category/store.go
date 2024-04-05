package category

import (
	"database/sql"
	"fmt"
	"log"
	"sparrow-plus/types"
	"strings"

	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) createCategoryTable() {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS "categories" (
		categoryId VARCHAR(36) PRIMARY KEY,
		name       VARCHAR(128) NOT NULL,
		createdAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deletedAt  TIMESTAMP
	)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Store) bulkInsertCategories(categories []string) {
	var insert = make([]string, 0, len(categories))
	var rowsArgs = make([]interface{}, 0, len(categories)*2)
	for _, categoryName := range categories {
		insert = append(insert, "(?, ?)")
		category := types.Category{
			CategoryId: uuid.New().String(),
			Name:       categoryName,
		}
		rowsArgs = append(rowsArgs, category.CategoryId)
		rowsArgs = append(rowsArgs, category.Name)

	}
	query := fmt.Sprintf(`
		INSERT INTO "categories"
			(categoryId, name)
		VALUES
			%v
	`, strings.Join(insert, ","))
	_, inerr := s.db.Exec(query, rowsArgs...)

	if inerr != nil {
		log.Fatal(inerr)
	}
}

func (s *Store) GetCategories() ([]*types.Category, error) {
	rows, err := s.db.Query(`SELECT * FROM "categories" WHERE "deletedAt" IS NULL`)

	if err != nil {
		return nil, err
	}

	categories := make([]*types.Category, 0)
	for rows.Next() {
		category, err := scanRowsIntoCategory(rows)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *Store) SetupCategories() {
	s.createCategoryTable()
	categories := []string{
		"TV Shows",
		"Movies",
		"Animes",
	}
	s.bulkInsertCategories(categories)
}

func scanRowsIntoCategory(rows *sql.Rows) (*types.Category, error) {
	category := new(types.Category)
	err := rows.Scan(
		&category.CategoryId,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return category, nil
}
