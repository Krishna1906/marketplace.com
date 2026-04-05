package rating

import (
	"encoding/json"
	"net/http"
	"strconv"

	"marketplace/internal/middleware"
)

// ➕ Add / Update Rating
func AddRatingHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req struct {
		ProductID int64 `json:"product_id"`
		Rating    int   `json:"rating"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := SaveRating(userID, req.ProductID, req.Rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Rating saved"))
}

// 📊 Get Rating
func GetRatingHandler(w http.ResponseWriter, r *http.Request) {

	productIDStr := r.URL.Query().Get("product_id")
	productID, _ := strconv.ParseInt(productIDStr, 10, 64)

	avg, count, err := FetchRating(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"average": avg,
		"count":   count,
	}

	json.NewEncoder(w).Encode(res)
}

func CanRateHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	productIDStr := r.URL.Query().Get("product_id")
	productID, _ := strconv.ParseInt(productIDStr, 10, 64)

	canRate, err := CanUserRate(userID, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{
		"can_rate": canRate,
	})
}