package order

type Item struct {
	Id    int     `json:"item-id"`
	Name  string  `json:"item-name"`
	Price float64 `json:"item-price"`
}

type Order struct {
	Id            int     `json:"order-id"`
	Items         []Item  `json:"items"`
	OrderStatus   string  `json:"order-status"`
	TotalPrice    float64 `json:"total-price"`
	PaymentStatus bool    `json:"payment-status"`
}
