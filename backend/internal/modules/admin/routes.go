package admin

import (
	"github.com/gorilla/mux"
	"marketplace/internal/middleware"
)

func RegisterAdminRoutes(router *mux.Router) {
	router.HandleFunc("/users/all", GetAllUsers).Methods("GET")
	router.Use(middleware.JWTAuthMiddleware)
	router.Use(middleware.RequireRole("ADMIN"))

	router.HandleFunc("/seller/approve", ApproveSellerHandler).Methods("POST")
	router.HandleFunc("/products/pending", GetPendingProducts).Methods("GET")
	router.HandleFunc("/product/action", ProductActionHandler).Methods("POST")
	router.HandleFunc("/products/all", GetAllProducts).Methods("GET")
	router.HandleFunc("/sellers/pending", GetPendingSellers).Methods("GET")
	router.HandleFunc("/sellers/all", GetAllSellers).Methods("GET")

}

