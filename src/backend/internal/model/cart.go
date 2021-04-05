package model

import "time"

type Cart struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	AddressID int `json:"address_id"`

	Token     string    `json:"token"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Content string `json:"content"`

	CartItems []*CartItem `json:"cart_items"`
}

type CartItem struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	CartID    int `json:"cart_id"`

	SKU      string  `json:"sku"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Quantity int     `json:"quantity"`

	Active    bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Content string `json:"content"`
}

func NewCart() *Cart {
	return &Cart{}
}
