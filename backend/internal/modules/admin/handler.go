package admin

import (
	"marketplace/internal/database"
	"marketplace/internal/utils"
	"net/http"
	"strconv"
)

func ApproveSellerHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	err = ApproveSeller(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Seller approved successfully"))
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, name, price, seller_id, status
		FROM products
		ORDER BY created_at DESC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.SellerID, &p.Status)
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	utils.JSON(w, http.StatusOK, products)
}
