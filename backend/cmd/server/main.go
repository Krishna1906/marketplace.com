package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"marketplace/internal/config"
	"marketplace/internal/database"
	"marketplace/internal/middleware"
	"marketplace/internal/modules/image"

	"marketplace/internal/modules/admin"
	"marketplace/internal/modules/auth"
	"marketplace/internal/modules/banner"
	"marketplace/internal/modules/cart"
	"marketplace/internal/modules/category"
	"marketplace/internal/modules/order"
	ordersummary "marketplace/internal/modules/order_summary"
	"marketplace/internal/modules/payment"
	"marketplace/internal/modules/product"
	"marketplace/internal/modules/seller"
	"marketplace/internal/modules/wishlist"
)

func main() {

	config.LoadConfig()
	database.Connect()

	r := mux.NewRouter()

	// Serve static files
	fs := http.FileServer(http.Dir("./uploads"))
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	// r := mux.NewRouter()
	r.Use(middleware.CORS)
	r.Use(middleware.Logger)

	api := r.PathPrefix("/api").Subrouter()

	// ---------- PUBLIC IMAGES ----------
	imageRouter := api.PathPrefix("/images").Subrouter()
	image.RegisterImageRoutes(imageRouter)
	api.Methods(http.MethodOptions)

	// ---------- AUTH (PUBLIC) ----------
	authRouter := api.PathPrefix("/auth").Subrouter()
	auth.RegisterAuthRoutes(authRouter)

	// ---------- PUBLIC ----------
	userCategoryRouter := api.PathPrefix("/categories").Subrouter()
	category.RegisterUserCategoryRoutes(userCategoryRouter)

	publicProductRouter := api.PathPrefix("/products").Subrouter()
	product.RegisterPublicProductRoutes(publicProductRouter)

	// ---------- SELLER APPLY (PUBLIC) ----------
	sellerRouter := api.PathPrefix("/seller").Subrouter()
	seller.RegisterSellerRoutes(sellerRouter)

	// ---------- SELLER PRODUCTS ----------
	sellerProductRouter := api.PathPrefix("/seller/products").Subrouter()
	sellerProductRouter.Use(middleware.JWTAuthMiddleware)
	sellerProductRouter.Use(middleware.RequireRole("SELLER"))
	product.RegisterSellerProductRoutes(sellerProductRouter)

	// ---------- ADMIN ----------
	adminRouter := api.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.JWTAuthMiddleware)
	adminRouter.Use(middleware.RequireRole("ADMIN"))

	admin.RegisterAdminRoutes(adminRouter)
	category.RegisterAdminCategoryRoutes(adminRouter)
	product.RegisterAdminProductRoutes(adminRouter)
	banner.RegisterBannerRoutes(r)

	// ---------- USER ----------
	cartRouter := api.PathPrefix("/cart").Subrouter()
	cartRouter.Use(middleware.JWTAuthMiddleware)
	cartRouter.Use(middleware.RequireRole("USER"))
	cart.RegisterCartRoutes(cartRouter)

	orderRouter := api.PathPrefix("/orders").Subrouter()
	orderRouter.Use(middleware.JWTAuthMiddleware)
	order.RegisterOrderRoutes(orderRouter)

	ordersummary.RegisterOrderSummaryRoutes(api)

	// ---------- PAYMENT ----------
	paymentRouter := api.PathPrefix("/payment").Subrouter()
	paymentRouter.Use(middleware.JWTAuthMiddleware)
	payment.RegisterPaymentRoutes(paymentRouter)

	// ---------- WISHLIST ----------
	// ---------- WISHLIST ----------
	wishlistRouter := api.PathPrefix("/wishlist").Subrouter()
	wishlistRouter.Use(middleware.JWTAuthMiddleware)
	wishlistRouter.Use(middleware.RequireRole("USER"))

	wishlist.RegisterWishlistRoutes(wishlistRouter)

	// ---------- HEALTH CHECK ----------

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	log.Println("🚀 Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
