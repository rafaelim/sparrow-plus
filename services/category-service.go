package services

import (
	"encoding/json"
	"net/http"
	customerrors "sparrow-plus/custom-errors"
	"sparrow-plus/db"
)


type CategoryService struct {
	dbHandlers *db.DB
}

func NewCategoryService(dbHandlers *db.DB) *CategoryService {
	return &CategoryService{
		dbHandlers: dbHandlers,
	}
}

func (cs *CategoryService) GetAll(w http.ResponseWriter, r *http.Request) error {
	categories, err := cs.dbHandlers.GetAllCategories()

	if err != nil {
		return customerrors.NewRequestError(
			http.StatusBadRequest,
			"Failed to get all categories",
			"FAILED_FIND_ALL",
		)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
	
	return nil
}
