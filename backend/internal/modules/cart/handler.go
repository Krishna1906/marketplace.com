package cart

import (
	"encoding/json"
	"net/http"

	"marketplace/internal/database"
	"marketplace/internal/middleware"

	"github.com/gorilla/mux"
)

type AddToCartRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

func AddToCart(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	var req AddToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO carts (user_id, product_id, quantity)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id, product_id)
	DO UPDATE SET quantity = carts.quantity + $3
	`

	_, err := database.DB.Exec(query, userID, req.ProductID, req.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product added to cart"))
}

func GetCart(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	query := `
	SELECT p.id, p.name, p.price, c.quantity
	FROM carts c
	JOIN products p ON p.id = c.product_id
	WHERE c.user_id = $1
	`

	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type CartItem struct {
		ProductID int64   `json:"product_id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		Quantity  int     `json:"quantity"`
	}

	var cart []CartItem

	for rows.Next() {
		var item CartItem
		rows.Scan(&item.ProductID, &item.Name, &item.Price, &item.Quantity)
		cart = append(cart, item)
	}

	json.NewEncoder(w).Encode(cart)
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int64)

	vars := mux.Vars(r)
	productID := vars["id"]

	query := `DELETE FROM carts WHERE user_id = $1 AND product_id = $2`

	result, err := database.DB.Exec(query, userID, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Item not found in cart", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product removed from cart"))
}

func UpdateCartQuantity(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)
	vars := mux.Vars(r)
	productID := vars["id"]

	type UpdateRequest struct {
		Quantity int `json:"quantity"`
	}

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result, err := UpdateCartService(userID, productID, req.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}