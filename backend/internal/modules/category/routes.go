package category

import (
	"github.com/gorilla/mux"
	"marketplace/internal/middleware"
)

func RegisterUserCategoryRoutes(router *mux.Router) {
	router.HandleFunc("", GetUserCategories).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", GetCategoryProducts).Methods("GET")

}

func RegisterSellerCategoryRoutes(router *mux.Router) {
	router.Use(middleware.JWTAuthMiddleware)
	router.Use(middleware.RequireRole("SELLER"))
	router.HandleFunc("", GetSellerCategories).Methods("GET")
}
func RegisterAdminCategoryRoutes(router *mux.Router) {
	// router is already /api/admin
	categoryRouter := router.PathPrefix("/categories").Subrouter()

	categoryRouter.Use(middleware.JWTAuthMiddleware)
	categoryRouter.Use(middleware.RequireRole("ADMIN"))

	categoryRouter.HandleFunc("", CreateCategory).Methods("POST")
	categoryRouter.HandleFunc("", GetAllCategories).Methods("GET")
	categoryRouter.HandleFunc("/{id:[0-9]+}", UpdateCategory).Methods("PUT")
	categoryRouter.HandleFunc("/{id:[0-9]+}", DeleteCategory).Methods("DELETE")
}