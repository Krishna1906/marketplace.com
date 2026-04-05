package ordersummary

type CartItem struct {
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	Image     string  `json:"image"`
}

type Address struct {
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	FullName    string `json:"full_name"`
	Phone       string `json:"phone"`
	AddressLine string `json:"address_line"`
	City        string `json:"city"`
	State       string `json:"state"`
	Pincode     string `json:"pincode"`
}

type OrderSummaryResponse struct {
	Items       []CartItem `json:"items"`
	TotalAmount float64    `json:"total_amount"`
	Addresses   []Address  `json:"addresses"`
}