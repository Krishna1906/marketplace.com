package cart

import "marketplace/internal/database"

func UpdateCartRepo(userID int64, productID string, qty int) (int, error) {

	// query := `
	// UPDATE carts
	// SET quantity = quantity + $1
	// WHERE user_id = $2 AND product_id = $3
	// RETURNING quantity
	// `

	query := `
	UPDATE carts
	SET quantity = $1
	WHERE user_id = $2 AND product_id = $3
	RETURNING quantity
	`

	var newQty int
	err := database.DB.QueryRow(query, qty, userID, productID).Scan(&newQty)

	if err != nil {
		return 0, err
	}

	return newQty, nil
}

func RemoveCartItemRepo(userID int64, productID string) error {

	query := `DELETE FROM carts WHERE user_id=$1 AND product_id=$2`

	_, err := database.DB.Exec(query, userID, productID)
	return err
}