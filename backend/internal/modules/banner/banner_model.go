package banner

type Banner struct {
	ID        int64  `json:"id"`
	ImageURL  string `json:"image_url"`
	ProductID int64  `json:"product_id"`
	IsActive  bool   `json:"is_active"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	CreatedAt string `json:"created_at"`
}