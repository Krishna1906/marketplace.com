package ordersummary

import (
	"github.com/gorilla/mux"
	"marketplace/internal/middleware"
)

func RegisterOrderSummaryRoutes(r *mux.Router) {
	router := r.PathPrefix("/ordersummary").Subrouter()

	router.Use(middleware.JWTAuthMiddleware)
	router.Use(middleware.RequireRole("USER"))

	router.HandleFunc("", GetOrderSummary).Methods("GET")
	router.HandleFunc("/address", SaveOrderAddress).Methods("POST")
	router.HandleFunc("/address/{id}", UpdateOrderAddress).Methods("PUT")
	router.HandleFunc("/address/{id}", DeleteOrderAddress).Methods("DELETE")
}