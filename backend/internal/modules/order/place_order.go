package order

import (
	"encoding/json"
	"net/http"

	"marketplace/internal/database"
	"marketplace/internal/middleware"
)
type PlaceOrderRequest struct {
	PaymentType string `json:"payment_type"` // COD | CARD | UPI
}

func PlaceOrderFromCart(w http.ResponseWriter, r *http.Request) {

	userIDVal := r.Context().Value(middleware.UserIDKey)
	userID, ok := userIDVal.(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// ✅ READ BODY
	var req PlaceOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ✅ VALIDATE
	if req.PaymentType == "" {
		http.Error(w, "payment_type is required", http.StatusBadRequest)
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// 1️⃣ Get cart items
	rows, err := tx.Query(`
		SELECT c.product_id, c.quantity, p.price, p.stock
		FROM carts c
		JOIN products p ON p.id = c.product_id
		WHERE c.user_id = $1
	`, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type CartItem struct {
		ProductID int64
		Quantity  int
		Price     float64
		Stock     int
	}

	var items []CartItem
	var total float64

	for rows.Next() {
		var item CartItem
		rows.Scan(&item.ProductID, &item.Quantity, &item.Price, &item.Stock)

		if item.Stock < item.Quantity {
			http.Error(w, "Insufficient stock", http.StatusBadRequest)
			return
		}

		total += item.Price * float64(item.Quantity)
		items = append(items, item)
	}

	if len(items) == 0 {
		http.Error(w, "Cart is empty", http.StatusBadRequest)
		return
	}

	// 2️⃣ ✅ CREATE ORDER WITH PAYMENT METHOD
	var orderID int64
	err = tx.QueryRow(`
		INSERT INTO orders (user_id, total_amount, payment_method, status)
		VALUES ($1, $2, $3, 'PLACED')
		RETURNING id
	`, userID, total, req.PaymentType).Scan(&orderID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3️⃣ Insert items + update stock
	for _, item := range items {

		_, err = tx.Exec(`
			INSERT INTO order_items (order_id, product_id, quantity, price)
			VALUES ($1, $2, $3, $4)
		`, orderID, item.ProductID, item.Quantity, item.Price)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = tx.Exec(`
			UPDATE products
			SET stock = stock - $1
			WHERE id = $2
		`, item.Quantity, item.ProductID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// 4️⃣ Clear cart
	_, err = tx.Exec(`DELETE FROM carts WHERE user_id = $1`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Commit failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":        "Order placed successfully",
		"order_id":       orderID,
		"total":          total,
		"payment_method": req.PaymentType,
	})
}