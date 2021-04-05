package model

import "time"

type Address struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`

	Default  bool   `json:"is_default"`
	Line1    string `json:"line_1"`
	Line2    string `json:"line_2"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
