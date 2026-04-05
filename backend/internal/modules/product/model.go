package product

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	SellerID    int64     `json:"seller_id"`
	CategoryID  int64     `json:"category_id"`
	CategoryName string `json:"category_name"`

	Category    string    `json:"category,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`

	Images []string `json:"images"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  int64   `json:"category_id"`
}
