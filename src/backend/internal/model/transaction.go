package model

import "time"

type Transaction struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	OrderID int    `json:"order_id"`
	Code    string `json:"code"`
	Type    int    `json:"type"`
	Mode    int    `json:"mode"`
	Status  int    `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Content string `json:"content"`
}
