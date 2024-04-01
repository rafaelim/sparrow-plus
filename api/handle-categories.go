package api

import (
	"net/http"
	"sparrow-plus/services"
)



func (s *APIServe) CategoriesHandlers(router *http.ServeMux) {
	categoryService := services.NewCategoryService(s.d)
	router.HandleFunc("GET /api/categories", withErrorHandling(categoryService.GetAll))
}
