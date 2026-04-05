package order

import (
	"database/sql"
	"marketplace/internal/database"
)

func GetProduct(productID int64) (float64, int, error) {
	var price float64
	var stock int

	err := database.DB.QueryRow(`
		SELECT price, stock FROM products WHERE id=$1
	`, productID).Scan(&price, &stock)

	return price, stock, err
}

func CreateOrderDB(tx *sql.Tx, userID int64, total float64, payment string) (int64, error) {
	var orderID int64
	err := tx.QueryRow(`
		INSERT INTO orders (user_id, total_amount, payment_method, status)
		VALUES ($1, $2, $3, 'PLACED')
		RETURNING id
	`, userID, total, payment).Scan(&orderID)

	return orderID, err
}

func InsertOrderItem(tx *sql.Tx, orderID, productID int64, qty int, price float64) error {
	_, err := tx.Exec(`
		INSERT INTO order_items (order_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
	`, orderID, productID, qty, price)

	return err
}