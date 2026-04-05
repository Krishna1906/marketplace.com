package seller

import (
	"github.com/gorilla/mux"
	"marketplace/internal/middleware"
)

func RegisterSellerRoutes(router *mux.Router) {
	// sellerRouter := router.PathPrefix("/seller").Subrouter()

	
	router.Use(middleware.JWTAuthMiddleware)
	router.Use(middleware.RequireRole("USER"))

	router.HandleFunc("/apply", ApplySeller).Methods("POST")
}
