package order

import (
	"net/http"
	"strconv"

	"marketplace/internal/database"
	"marketplace/internal/middleware"

	"github.com/gorilla/mux"
)

func CancelOrder(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	params := mux.Vars(r)
	orderID, _ := strconv.ParseInt(params["id"], 10, 64)

	// ❌ Don't cancel if already shipped/delivered
	var status string
	err := database.DB.QueryRow(`
		SELECT status FROM orders WHERE id=$1 AND user_id=$2
	`, orderID, userID).Scan(&status)

	if err != nil {
		http.Error(w, "Order not found", 404)
		return
	}

	if status != "PLACED" {
		http.Error(w, "Cannot cancel this order", 400)
		return
	}

	_, err = database.DB.Exec(`
		UPDATE orders SET status='CANCELLED' WHERE id=$1
	`, orderID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("Order cancelled successfully"))
}