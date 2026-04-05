package admin

import (
	"net/http"

	"marketplace/internal/database"
	"marketplace/internal/utils"
)

type PendingSellerResponse struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	ShopName  string `json:"shop_name"`
	GSTNumber string `json:"gst_number"`
	Status    string `json:"status"`
}

func GetPendingSellers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, user_id, shop_name, gst_number, status
		FROM sellers
		WHERE status = 'PENDING'
		ORDER BY id DESC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sellers []PendingSellerResponse

	for rows.Next() {
		var s PendingSellerResponse
		if err := rows.Scan(
			&s.ID,
			&s.UserID,
			&s.ShopName,
			&s.GSTNumber,
			&s.Status,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sellers = append(sellers, s)
	}

	utils.JSON(w, http.StatusOK, sellers)
}

func GetAllSellers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, user_id, shop_name, gst_number, status
		FROM sellers
		ORDER BY id DESC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sellers []PendingSellerResponse

	for rows.Next() {
		var s PendingSellerResponse
		if err := rows.Scan(
			&s.ID,
			&s.UserID,
			&s.ShopName,
			&s.GSTNumber,
			&s.Status,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sellers = append(sellers, s)
	}

	utils.JSON(w, http.StatusOK, sellers)
}
