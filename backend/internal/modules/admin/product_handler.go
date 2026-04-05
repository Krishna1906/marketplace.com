package admin

import (
	// "encoding/json"
	"marketplace/internal/database"
	"marketplace/internal/utils"
	"net/http"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	SellerID    int64   `json:"seller_id"`
	Status      string  `json:"status"`
}

func GetPendingProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, name, price, seller_id, status
		FROM products
		WHERE status = 'PENDING'
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
