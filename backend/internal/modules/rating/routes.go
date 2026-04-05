package rating

import "github.com/gorilla/mux"

func RegisterRatingRoutes(router *mux.Router) {

	router.HandleFunc("", AddRatingHandler).Methods("POST")
	router.HandleFunc("", GetRatingHandler).Methods("GET")

	// 🔥 ADD THIS (you forgot probably)
	router.HandleFunc("/can-rate", CanRateHandler).Methods("GET")
}