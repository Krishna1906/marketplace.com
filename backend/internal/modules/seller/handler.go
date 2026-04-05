package seller

import (
	"encoding/json"
	"net/http"

	"marketplace/internal/database"
	"marketplace/internal/middleware"
)

type SellerApplication struct {
	ShopName  string `json:"shop_name"`
	GSTNumber string `json:"gst_number"`
}

func ApplySeller(w http.ResponseWriter, r *http.Request) {

	userIDVal := r.Context().Value(middleware.UserIDKey)
	userID, ok := userIDVal.(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req SellerApplication
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.ShopName == "" {
		http.Error(w, "shop_name is required", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO sellers (user_id, shop_name, gst_number, status)
		VALUES ($1, $2, $3, 'PENDING')
		ON CONFLICT (user_id) DO NOTHING
	`

	_, err := database.DB.Exec(
		query,
		userID,
		req.ShopName,
		req.GSTNumber,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Seller application submitted"))
}
func ApproveSeller(w http.ResponseWriter, r *http.Request) {

	userIDVal := r.Context().Value(middleware.UserIDKey)
	userID, ok := userIDVal.(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	query := `
		UPDATE sellers
		SET status = $1
		WHERE user_id = $2
	`

	_, err := database.DB.Exec(query, "APPROVED", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Seller approved"))
}