package image

import "time"

type Image struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}