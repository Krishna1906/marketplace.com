package routes

import (
	"database/sql"
	"marketplace/internal/modules/admin"
	"marketplace/internal/modules/auth"
	"marketplace/internal/modules/cart"
	"marketplace/internal/modules/image"
	"marketplace/internal/modules/order"

	// "marketplace/internal/modules/product"
	"marketplace/internal/modules/seller"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all module routes for the application
func RegisterRoutes(r *mux.Router, db *sql.DB) {
	// API base path
	api := r.PathPrefix("/api").Subrouter()

	admin.RegisterAdminRoutes(api)
	auth.RegisterAuthRoutes(api)
	cart.RegisterCartRoutes(api)
	order.RegisterOrderRoutes(api)
	// product.RegisterProductRoutes(api)
	seller.RegisterSellerRoutes(api)	

	image.RegisterImageRoutes(api)	
}
