package store

import "github.com/iftech-a/lookum/src/backend/internal/model"

//ProductRepo repository for working with database on products
type ProductRepo interface {
	Create(*model.Product) error
	GetProducts(int, int) ([]*model.Product, error)
	GetProduct(int) (*model.Product, error)
	AddImage(int, string, string) error
}
