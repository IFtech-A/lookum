package model

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//Product defines product's attributes
type Product struct {
	ID        int     `json:"id" path:"id"`
	UserID    int     `json:"vendor_id"`
	Title     string  `json:"title"`
	MetaTitle string  `json:"meta_title"`
	Slug      string  `json:"slug"`
	Summary   string  `json:"summary"`
	Type      int     `json:"type"`
	SKU       string  `json:"sku"`
	Price     float32 `json:"price"`
	Discount  float32 `json:"discount"`
	Quantity  int     `json:"quantity"`
	Available bool    `json:"shop_available"`
	Content   string  `json:"content"`

	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`

	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`

	Images []*Image `json:"images,omitempty"`
}

//Image structure defines image attributes
type Image struct {
	ID       int    `json:"id"`
	FileURI  string `json:"file_uri"`
	Filename string `json:"-"`
	Main     bool   `json:"main"`
}

//NewProduct ...
func NewProduct() *Product {
	return &Product{}
}

func (p *Product) GenerateSlug() {
	hash := sha256.Sum256([]byte(p.Title))
	p.Slug = hex.EncodeToString(hash[:])
}
