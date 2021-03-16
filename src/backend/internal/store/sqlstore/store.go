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
