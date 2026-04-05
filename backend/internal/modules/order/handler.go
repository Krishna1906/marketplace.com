package order

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"marketplace/internal/database"
	"marketplace/internal/middleware"
	"marketplace/internal/utils"

	"github.com/gorilla/mux"
)

// =====================
// CREATE ORDER
// =====================
func CreateOrder(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	var req CreateOrderRequest
	json.NewDecoder(r.Body).Decode(&req)

	if err := ValidatePayment(req.PaymentType); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	tx, _ := database.DB.Begin()
	defer tx.Rollback()

	price, stock, err := GetProduct(req.ProductID)
	if err != nil {
		http.Error(w, "Product not found", 404)
		return
	}

	if stock < req.Quantity {
		http.Error(w, "Insufficient stock", 400)
		return
	}

	total := CalculateTotal(price, req.Quantity)

	// update stock
	tx.Exec(`UPDATE products SET stock = stock - $1 WHERE id=$2`,
		req.Quantity, req.ProductID)

	orderID, err := CreateOrderDB(tx, userID, total, req.PaymentType)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	InsertOrderItem(tx, orderID, req.ProductID, req.Quantity, price)

	tx.Commit()

	utils.JSON(w, 201, map[string]interface{}{
		"order_id": orderID,
		"total":    total,
	})
}

// =====================
// GET ALL ORDERS
// =====================
func GetOrders(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	rows, _ := database.DB.Query(`
		SELECT id, total_amount, payment_method, status, created_at
		FROM orders WHERE user_id=$1 ORDER BY created_at DESC
	`, userID)

	defer rows.Close()

	var orders []OrderResponse

	for rows.Next() {

		var o OrderResponse
		var payment sql.NullString

		rows.Scan(&o.OrderID, &o.TotalAmount, &payment, &o.Status, &o.CreatedAt)

		if payment.Valid {
			o.PaymentMethod = payment.String
		} else {
			o.PaymentMethod = "UNKNOWN"
		}

		itemRows, _ := database.DB.Query(`
			SELECT oi.product_id, p.name, oi.price, oi.quantity,
			COALESCE((SELECT image_url FROM product_images WHERE product_id=p.id LIMIT 1),'')
			FROM order_items oi
			JOIN products p ON p.id = oi.product_id
			WHERE oi.order_id=$1
		`, o.OrderID)

		for itemRows.Next() {
			var item OrderItem
			itemRows.Scan(&item.ProductID, &item.Name, &item.Price, &item.Quantity, &item.Image)
			o.Items = append(o.Items, item)
		}
		itemRows.Close()

		orders = append(orders, o)
	}

	utils.JSON(w, 200, orders)
}

func TrackOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int64)

	orderIDParam := mux.Vars(r)["order_id"]
	orderID, _ := strconv.ParseInt(orderIDParam, 10, 64)

	var status string
	var createdAt string

	err := database.DB.QueryRow(`
		SELECT status, created_at
		FROM orders
		WHERE id = $1 AND user_id = $2
	`, orderID, userID).Scan(&status, &createdAt)

	if err != nil {
		http.Error(w, "Order not found", 404)
		return
	}

	utils.JSON(w, http.StatusOK, map[string]interface{}{
		"order_id":  orderID,
		"status":    status,
		"created_at": createdAt,
	})
}