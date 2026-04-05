package product

import "github.com/gorilla/mux"

// ---------- PUBLIC ----------
// Duplicate function removed to avoid redeclaration error.

// ---------- SELLER ----------
func RegisterSellerProductRoutes(router *mux.Router) {
	// /api/seller/products
	router.HandleFunc("", CreateProduct).Methods("POST")
	router.HandleFunc("", GetProducts).Methods("GET")

	// /api/seller/products/{id}/images
	router.HandleFunc("/{id:[0-9]+}/images", UploadProductImages).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}/update", UpdateProduct).Methods("PUT")
	router.HandleFunc("/{id:[0-9]+}", DeleteProduct).Methods("DELETE")
}


func RegisterPublicProductRoutes(router *mux.Router) {
	router.HandleFunc("", GetProducts).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", GetProductByID).Methods("GET")
}


