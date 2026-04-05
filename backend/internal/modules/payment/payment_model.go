package payment

import "time"

type CardPayment struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CardLast4 string    `json:"card_last4"`
	CreatedAt time.Time `json:"created_at"`
}
