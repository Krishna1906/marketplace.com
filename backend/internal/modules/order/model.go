package order

type CreateOrderRequest struct {
	ProductID   int64  `json:"product_id"`
	Quantity    int    `json:"quantity"`
	PaymentType string `json:"payment_type"`
}

type OrderItem struct {
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Image     string  `json:"image"`
}

type OrderResponse struct {
	OrderID       int64       `json:"order_id"`
	TotalAmount   float64     `json:"total_amount"`
	PaymentMethod string      `json:"payment_method"`
	Status        string      `json:"status"`
	CreatedAt     string      `json:"created_at"`
	Items         []OrderItem `json:"items"`
}