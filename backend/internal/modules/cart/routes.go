package cart

import "github.com/gorilla/mux"

func RegisterCartRoutes(router *mux.Router) {
	router.HandleFunc("", AddToCart).Methods("POST")
	router.HandleFunc("", GetCart).Methods("GET")
	router.HandleFunc("/{id}", RemoveFromCart).Methods("DELETE")
	router.HandleFunc("/update/{id}", UpdateCartQuantity).Methods("PUT")
}
