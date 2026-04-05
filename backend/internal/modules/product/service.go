package product

import (
	"errors"
	"marketplace/internal/database"
	"marketplace/internal/modules/seller"
)

func ValidateSeller(userID int64) error {
	return seller.EnsureSellerApproved(userID)
}

func AddProduct(sellerID int64, p *Product) error {
	if sellerID == 0 {
		return errors.New("unauthorized seller")
	}

	if p.Name == "" || p.Price <= 0 {
		return errors.New("invalid product data")
	}

	p.SellerID = sellerID
	p.Status = "DRAFT"

	return InsertProduct(p)
}

func ApproveSeller(userID int64) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE sellers
		SET status = 'APPROVED'
		WHERE user_id = $1
	`, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		UPDATE products
		SET status = 'APPROVED'
		WHERE seller_id = $1
	`, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
