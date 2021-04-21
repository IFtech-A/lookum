package store

import "github.com/iftech-a/lookum/src/backend/internal/model"

//ProductRepo repository for working with database on products
type ProductRepo interface {
	Create(*model.Product) error
	GetProducts(int, int) ([]*model.Product, error)
	GetProduct(int) (*model.Product, error)
	UpdateProduct(*model.Product) error
	DeleteProduct(int) error
	AddImage(int, string, string) error
}

//OrderRepo repository for working with database on product orders
type OrderRepo interface {
	Create(*model.Order) (int, error)
	GetOrders(int, int) ([]*model.Order, error)
	GetOrder(int) (*model.Order, error)
	GetOrderWithItems(int) (*model.Order, error)
	DeleteOrder(int) error
}

//CartRepo repository for working with database on product orders
type CartRepo interface {
	Create(*model.Cart) (int, error)
	GetCarts(int, int) ([]*model.Cart, error)
	GetCart(int) (*model.Cart, error)
	GetCartWithItems(int) (*model.Cart, error)
	DeleteCart(int) error
}

//UserRepo repository for working with database on product orders
type UserRepo interface {
	Create(*model.User) (int, error)
	GetUser(int) (*model.User, error)
	GetUserByEmail(string) (*model.User, error)
	DeleteUser(int) error
}
