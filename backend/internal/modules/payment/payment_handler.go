package payment

import (
	"net/http"
	"encoding/json"

	"marketplace/internal/utils"
	"marketplace/internal/database"
	"marketplace/internal/middleware"
)

func GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, map[string]interface{}{
		"methods": []string{
			"CASH_ON_DELIVERY",
			"CARD",
			"UPI",
		},
	})
}

type CardRequest struct {
	CardNumber string `json:"card_number"`
	Expiry     string `json:"expiry"`
	CVV        string `json:"cvv"`
}

func SaveCardPayment(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	var req CardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if len(req.CardNumber) < 12 {
		http.Error(w, "Invalid card number", http.StatusBadRequest)
		return
	}

	if len(req.CVV) < 3 {
		http.Error(w, "Invalid CVV", http.StatusBadRequest)
		return
	}

	// ✅ STORE ONLY LAST 4 DIGITS
	last4 := req.CardNumber[len(req.CardNumber)-4:]

	_, err := database.DB.Exec(`
		INSERT INTO payments (user_id, method, card_last4, expiry)
		VALUES ($1, 'CARD', $2, $3)
	`, userID, last4, req.Expiry)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusCreated, map[string]string{
		"message": "Card saved securely",
	})
}