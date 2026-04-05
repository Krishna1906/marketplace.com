package product

import "marketplace/internal/database"

func InsertProduct(p *Product) error {
	query := `
		INSERT INTO products
		(seller_id, name, description, price, stock, category_id, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := database.DB.Exec(
		query,
		p.SellerID,
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		p.CategoryID, // ✅ FIX
		p.Status,
	)
	return err
}
