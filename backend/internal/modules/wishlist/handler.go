package wishlist

import (
	"encoding/json"
	"net/http"

	"marketplace/internal/database"
	"marketplace/internal/middleware"
)

// ======================
// ➜ ADD TO WISHLIST
// ======================
type AddWishlistRequest struct {
	ProductID int64 `json:"product_id"`
}

func AddToWishlist(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
if !ok {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
}

	var req AddWishlistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO wishlist (user_id, product_id)
	VALUES ($1, $2)
	ON CONFLICT (user_id, product_id) DO NOTHING
	`

	_, err := database.DB.Exec(query, userID, req.ProductID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Added to wishlist"))
}

// ======================
// ➜ GET WISHLIST
// ======================
func GetWishlist(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	query := `
	SELECT DISTINCT ON (p.id) 
		p.id, 
		p.name, 
		p.price, 
		pi.image_url
	FROM wishlist w
	JOIN products p ON p.id = w.product_id
	LEFT JOIN product_images pi ON pi.product_id = p.id
	WHERE w.user_id = $1
	ORDER BY p.id, pi.id ASC
	`

	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Item struct {
		ProductID int64   `json:"product_id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		Image     *string `json:"image"`
	}

	var list []Item

	for rows.Next() {
		var i Item

		err := rows.Scan(&i.ProductID, &i.Name, &i.Price, &i.Image)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		list = append(list, i)
	}

	if list == nil {
		list = []Item{}
	}

	json.NewEncoder(w).Encode(list)
}
// ======================
// ➜ REMOVE FROM WISHLIST
// ======================
func RemoveFromWishlist(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
if !ok {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
}
	productID := r.URL.Query().Get("product_id")

	query := `DELETE FROM wishlist WHERE user_id=$1 AND product_id=$2`

	result, err := database.DB.Exec(query, userID, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Removed from wishlist"))
}

// ======================
// 🔥 MOVE TO CART (MAIN FEATURE)
// ======================
func MoveToCart(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
if !ok {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
}

	var req struct {
		ProductID int64 `json:"product_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 1️⃣ Add to cart (same logic as your AddToCart)
	cartQuery := `
	INSERT INTO carts (user_id, product_id, quantity)
	VALUES ($1, $2, 1)
	ON CONFLICT (user_id, product_id)
	DO UPDATE SET quantity = carts.quantity + 1
	`

	_, err := database.DB.Exec(cartQuery, userID, req.ProductID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2️⃣ Remove from wishlist
	_, err = database.DB.Exec(
		`DELETE FROM wishlist WHERE user_id=$1 AND product_id=$2`,
		userID,
		req.ProductID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Moved to cart successfully"))
}