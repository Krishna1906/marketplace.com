package image

import (
	"marketplace/internal/database"

	"github.com/gorilla/mux"
)

func RegisterImageRoutes(r *mux.Router) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	handler := NewHandler(service)

	// 🔥 THIS IS YOUR NEW API
	r.HandleFunc("/product/{id}", handler.GetImagesByProductID).Methods("GET")
}