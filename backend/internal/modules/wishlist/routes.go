package wishlist

import "github.com/gorilla/mux"

func RegisterWishlistRoutes(router *mux.Router) {

	router.HandleFunc("", AddToWishlist).Methods("POST")
	router.HandleFunc("", GetWishlist).Methods("GET")
	router.HandleFunc("/remove", RemoveFromWishlist).Methods("DELETE")

	// 🔥 MAIN ROUTE
	router.HandleFunc("/move-to-cart", MoveToCart).Methods("POST")
}