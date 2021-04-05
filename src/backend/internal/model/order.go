package model

import "time"

const (
	OrderAccepted = iota + 1000
	OrderDeclined
	OrderCanceled
	OrderDelayed
	OrderDelivering
	OrderCompleted
)

type Order struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	AddressID int    `json:"address_id"`
	Token     string `json:"token"`
	Status    int    `json:"status"`

	SubTotal     float64 `json:"sub_total"`
	ItemDiscount float64 `json:"item_discount"`
	Tax          float32 `json:"tax"`
	Shipping     float32 `json:"shipping"`
	Total        float64 `json:"total"`

	Promo         string  `json:"promo"`
	TotalDiscount float32 `json:"total_discount"`
	GrandTotal    float64 `json:"grand_total"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	OrderItems []*OrderItem `json:"order_items,omitempty"`
}

type OrderItem struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	OrderID   int `json:"order_id"`

	SKU      string  `json:"sku"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Quantity float32 `json:"quantity"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOrder() *Order {
	return &Order{
		Status:    OrderAccepted,
		CreatedAt: time.Now().In(time.UTC),
	}
}

func NewOrderItem(productID int, quantity, price float32) *OrderItem {
	return &OrderItem{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
}
