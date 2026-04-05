package image

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetByProductID(productID int) ([]Image, error) {
	query := `
		SELECT id, product_id, image_url, created_at
		FROM product_images
		WHERE product_id = $1
		ORDER BY id ASC
	`

	rows, err := r.DB.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []Image

	for rows.Next() {
		var img Image
		err := rows.Scan(&img.ID, &img.ProductID, &img.ImageURL, &img.CreatedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	return images, nil
}