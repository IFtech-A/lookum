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
	ID         int          `json:"id"`
	UserID     int          `json:"user_id"`
	Status     int          `json:"status"`
	CreatedAt  time.Time    `json:"created_at"`
	OrderItems []*OrderItem `json:"order_items,omitempty"`
}

type OrderItem struct {
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  float32 `json:"quantity"`
	AtPrice   float32 `json:"at_price"`
}

func NewOrder() *Order {
	return &Order{
		Status:    OrderAccepted,
		CreatedAt: time.Now().In(time.UTC),
	}
}

func NewOrderItem(productID int, quantity, atPrice float32) *OrderItem {
	return &OrderItem{
		ProductID: productID,
		Quantity:  quantity,
		AtPrice:   atPrice,
	}
}
