package rating

import "marketplace/internal/database"

// ➕ Add or Update Rating
func AddOrUpdateRating(userID, productID int64, rating int) error {
	query := `
	INSERT INTO ratings (user_id, product_id, rating)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id, product_id)
	DO UPDATE SET rating = EXCLUDED.rating
	`

	_, err := database.DB.Exec(query, userID, productID, rating)
	return err
}

// 📊 Get Average Rating
func GetProductRating(productID int64) (float64, int, error) {
	query := `
	SELECT COALESCE(AVG(rating), 0), COUNT(*)
	FROM ratings
	WHERE product_id = $1
	`

	var avg float64
	var count int

	err := database.DB.QueryRow(query, productID).Scan(&avg, &count)
	return avg, count, err
}

func CanUserRate(userID, productID int64) (bool, error) {
	query := `
	SELECT COUNT(*) 
	FROM order_items oi
	JOIN orders o ON o.id = oi.order_id
	WHERE o.user_id = $1 
	AND oi.product_id = $2 
	AND oi.status = 'DELIVERED'
	`

	var count int
	err := database.DB.QueryRow(query, userID, productID).Scan(&count)
	return count > 0, err
}