package store

import "lookum/internal/model"

//ProductRepo repository for working with database on products
type ProductRepo interface {
	Create(*model.Product) error
	FindByID(int) (*model.Product, error)
	GetProducts(int, int) ([]*model.Product, error)
	GetProduct(int) (*model.Product, error)
	AddImage(int, string, string) error
}
