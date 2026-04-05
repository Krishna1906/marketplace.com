package banner

import "github.com/gorilla/mux"

func RegisterBannerRoutes(r *mux.Router) {

	// 🔐 ADMIN
	admin := r.PathPrefix("/api/admin/banner").Subrouter()
	admin.HandleFunc("/upload", UploadBanner).Methods("POST")
	admin.HandleFunc("", GetAdminBanners).Methods("GET")
	// ✅ NEW UPDATE API
	r.HandleFunc("/api/admin/banner/{id}/update", UpdateBanner).Methods("POST")

	// 🌍 USER
	r.HandleFunc("/api/banner", GetUserBanners).Methods("GET")
}