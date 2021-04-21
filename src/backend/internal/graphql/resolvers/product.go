package resolver

import (
	"errors"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/iftech-a/lookum/src/backend/internal/store"
)

func Product(s store.ProductRepo) func(graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		idStr, ok := p.Args["id"]
		if !ok {
			return nil, errors.New("product id parameter missing")
		}

		id, err := strconv.Atoi(idStr.(string))
		if err != nil {
			return nil, errors.New("malformed param")
		}

		return s.GetProduct(id)
	}
}

func Products(s store.ProductRepo) func(graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {

		categoryId := 0
		idStr, ok := p.Args["category_id"]
		if ok {
			id, err := strconv.Atoi(idStr.(string))
			if err == nil {
				categoryId = id
			}
		}

		limit := p.Args["limit"].(int)

		return s.GetProducts(limit, categoryId)
	}
}
