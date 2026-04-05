package product

import (
	"marketplace/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterAdminProductRoutes(router *mux.Router) {

	admin := router.PathPrefix("/admin/products").Subrouter()
	admin.Use(middleware.JWTAuthMiddleware)
	admin.Use(middleware.RequireRole("ADMIN"))

	// GET /api/admin/products/all
	admin.HandleFunc("/all", AdminGetAllProducts).Methods("GET")
}
