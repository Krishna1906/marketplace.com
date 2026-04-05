package cart

type Cart struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	ProductID uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}