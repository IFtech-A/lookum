package sqlstore

import (
	"database/sql"

	"github.com/iftech-a/lookum/src/backend/internal/store"

	_ "github.com/lib/pq" // postgresql pq library
)

//Store ...
type Store struct {
	db          *sql.DB
	productRepo *ProductRepo
	orderRepo   *OrderRepo
	cartRepo    *CartRepo
	userRepo    *UserRepo
}

//New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//Product returns repository with Product related API
func (s *Store) Product() store.ProductRepo {
	if s.productRepo == nil {
		s.productRepo = &ProductRepo{
			store: s,
		}
	}

	return s.productRepo
}

//Order returns repository with Order related API
func (s *Store) Order() store.OrderRepo {
	if s.orderRepo == nil {
		s.orderRepo = &OrderRepo{
			store: s,
		}
	}

	return s.orderRepo
}

//Cart returns repository with Cart related API
func (s *Store) Cart() store.CartRepo {
	if s.cartRepo == nil {
		s.cartRepo = &CartRepo{
			store: s,
		}
	}

	return s.cartRepo
}

//User returns repository with User related API
func (s *Store) User() store.UserRepo {
	if s.userRepo == nil {
		s.userRepo = &UserRepo{
			store: s,
		}
	}

	return s.userRepo
}
