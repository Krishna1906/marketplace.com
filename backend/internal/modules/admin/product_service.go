package admin

import (
	"errors"
	"marketplace/internal/database"
)

func UpdateProductStatus(productID int64, status string) error {
	if status != "APPROVED" && status != "REJECTED" {
		return errors.New("invalid status")
	}

	result, err := database.DB.Exec(`
		UPDATE products
		SET status = $1
		WHERE id = $2
	`, status, productID)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}
