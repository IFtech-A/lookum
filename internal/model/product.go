package model

import "time"

//Product defines product's attributes
type Product struct {
	ID          int       `json:"id" path:"id" query:"id" form:"id"`
	CategoryID  int       `json:"category_id" path:"category_id" query:"category_id" form:"category_id"`
	Name        string    `json:"title" path:"name" query:"name" form:"name"`
	Description string    `json:"desc" path:"desc" query:"desc" form:"desc"`
	Price       float32   `json:"price" path:"price" query:"price" form:"price"`
	Discount    float32   `json:"discount" path:"discount" query:"discount" form:"discount"`
	Status      string    `json:"stock_status"`
	Stock       int       `json:"in_stock" path:"in_stock" query:"in_stock" form:"in_stock"`
	Likes       int       `json:"likes"`
	CreatedAt   time.Time `json:"created_at"`
}

//NewProduct ...
func NewProduct() *Product {
	return &Product{}
}
