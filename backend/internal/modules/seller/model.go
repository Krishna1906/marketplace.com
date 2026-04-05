package seller

type Seller struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	ShopName  string `json:"shop_name"`
	GSTNumber string `json:"gst_number"`
	Status    string `json:"status"`
}
