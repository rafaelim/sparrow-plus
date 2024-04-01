package db

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	CategoryId string `json:"categoryId"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func NewCategory(name string) Category {
	return Category{
		CategoryId: uuid.New().String(),
		Name: name,
	}
}

func (d * DB) createCategoryTable () {
	_, err := d.dbInstance.Exec(`
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

func (d *DB) bulkInsertCategories(categories []string) {
	var insert = make([]string, 0, len(categories))
	var rowsArgs = make([]interface{}, 0, len(categories) * 2)
	for _, v := range categories {
		insert = append(insert, "(?, ?)")
		category := NewCategory(v)
		rowsArgs = append(rowsArgs, category.CategoryId)
		rowsArgs = append(rowsArgs, category.Name)
		
	}
	query := fmt.Sprintf(`
		INSERT INTO "categories"
			(categoryId, name)
		VALUES
			%v
	`, strings.Join(insert, ","))
	_, inerr := d.dbInstance.Exec(query, rowsArgs...)
	
	
	if inerr != nil {
		log.Fatal(inerr)
	}
}

	
func (d *DB) GetAllCategories() ([]Category, error) {
	rows, err := d.dbInstance.Query(`SELECT * FROM "categories" WHERE "deletedAt" IS NULL`)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.CategoryId,
			&category.Name, 
			&category.CreatedAt, 
			&category.UpdatedAt,
			&category.DeletedAt,
		)

		if err != nil {
			log.Println(err.Error())
			return nil, fmt.Errorf("failed to scan categories")
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (d *DB) SetupCategories() {
	d.createCategoryTable()
	categories := []string{
		"TV Shows",
		"Movies",
		"Animes",
	}
	d.bulkInsertCategories(categories)
}