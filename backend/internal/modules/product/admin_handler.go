package product

import (
	"net/http"

	"marketplace/internal/database"
	"marketplace/internal/utils"
)

type AdminProductResponse struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Status      string   `json:"status"`
	SellerID    int64    `json:"seller_id"`
	SellerName  string   `json:"seller_name"`
	CategoryID  int64    `json:"category_id"`
	Category    string   `json:"category"`
	CreatedAt   string   `json:"created_at"`
	Images      []string `json:"images"`
}

func AdminGetAllProducts(w http.ResponseWriter, r *http.Request) {

	query := `
	SELECT 
		p.id, p.name, p.description, p.price, p.stock,
		p.status, p.seller_id,
		u.name AS seller_name,
		p.category_id, c.name AS category_name,
		p.created_at
	FROM products p
	JOIN users u ON u.id = p.seller_id
	JOIN categories c ON c.id = p.category_id
	ORDER BY p.created_at DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []AdminProductResponse

	for rows.Next() {
		var p AdminProductResponse

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.Status,
			&p.SellerID,
			&p.SellerName,
			&p.CategoryID,
			&p.Category,
			&p.CreatedAt,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p.Images, _ = GetProductImages(p.ID)
		products = append(products, p)
	}

	utils.JSON(w, http.StatusOK, products)
}
