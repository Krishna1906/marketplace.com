package rating

type Rating struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	Rating    int   `json:"rating"`
}